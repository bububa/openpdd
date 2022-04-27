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

// Read read pmc message
func Read(ctx context.Context, clt *core.SDKClient) <-chan Command {
	ch := make(chan Command)
	go func() {
		defer close(ch)
		conn, err := reconnect(ctx, clt)
		if err != nil {
			return
		}
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
						ch <- cmd
					}
				}
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
