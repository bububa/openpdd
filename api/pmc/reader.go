package pmc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/bububa/openpdd/core"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func Read(ctx context.Context, clt *core.SDKClient) <-chan struct {
	Command
	HandleResult
} {
	ch := make(chan struct {
		Command
		HandleResult
	})
	conn, err := reconnect(ctx, clt)
	if err != nil {
		fmt.Println(err)
	}
	result := make(HandleResult)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if payload, err := wsutil.ReadServerText(conn); err != nil {
					time.Sleep(1 * time.Second)
					if conn, err = reconnect(ctx, clt); err != nil {
						return
					}
				} else {
					var cmd Command
					if err := json.Unmarshal(payload, &cmd); err == nil {
						ch <- struct {
							Command
							HandleResult
						}{Command: cmd, HandleResult: result}

						select {
						case err := <-result:
							if err == nil {
								err = NewAckMessage(cmd.ID, cmd.SendTime, cmd.Message.Type, cmd.Message.MallID).
									sendAck(conn)
							}
							if err != nil {
								fmt.Println(err)
							}

						}
					}
				}
			}
		}
	}()

	go func() {
		interval := time.NewTicker(5 * time.Second)
		for range interval.C {
			err := NewHeartBeatMessage().sendHeartBeat(conn)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	return ch
}

func reconnect(ctx context.Context, clt *core.SDKClient) (net.Conn, error) {
	conn, _, _, err := ws.Dial(ctx, clt.WSSUrl())
	if err != nil {
		fmt.Println(err)
	}
	return conn, err
}
