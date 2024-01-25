package pmc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobwas/ws/wsutil"
	"net"
	"time"
)

type CommandType string

const (
	CommandType_COMMON    CommandType = "Common"
	CommandType_FAIL      CommandType = "Fail"
	CommandType_HEARTBEAT CommandType = "HeartBeat"
	CommandType_ACK       CommandType = "Ack"
)

type EventType string

const (
	EventType_TradeModifiedEvent EventType = "pdd_ddjb_TradeModified"
)

// Command pmc record
type Command struct {
	ID          uint64      `json:"id,omitempty"`
	CommandType CommandType `json:"commandType,omitempty"`
	Time        int64       `json:"time,omitempty"`
	SendTime    int64       `json:"sendTime,omitempty"`
	Message     Message     `json:"message,omitempty"`
}

func (c Command) IsError() bool {
	return c.CommandType == CommandType_FAIL
}

func (c Command) Error() string {
	return c.Message.Content
}

func (c Command) Event() (Event, error) {
	if c.IsError() {
		return nil, c
	}
	switch c.Message.Type {
	case EventType_TradeModifiedEvent:
		var ev TradeModifiedEvent
		if err := json.Unmarshal([]byte(c.Message.Content), &ev); err != nil {
			fmt.Println(c.Message.Content)
			return nil, err
		}
		return &ev, nil
	}
	return nil, errors.New("unknown event")
}

type Message struct {
	// Type event type
	Type EventType `json:"type,omitempty"`
	// Content
	Content string `json:"content,omitempty"`
	// MallID 商户ID
	MallID int64 `json:"mallID,omitempty"`
}

// Event pmc event
type Event interface {
	Type() EventType
}

// TradeModifiedEvent 多多进宝订单状态变更
type TradeModifiedEvent struct {
	// Tid 订单号
	Tid string `json:"tid,omitempty"`
	// Status 订单状态：0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算 ;10-已处罚
	Status int `json:"status,omitempty"`
	// Pid 推广位ID
	Pid string `json:"pid,omitempty"`
	// CustomParameters 代理身份自定义参数
	CustomParameters string `json:"custom_parameters,omitempty"`
	// ModifyTime 最后更新时间
	ModifyTime int64 `json:"modify_time,omitempty"`
}

// Type implement Event interface
func (m TradeModifiedEvent) Type() EventType {
	return EventType_TradeModifiedEvent
}

type AckMessage struct {
	ID          uint64      `json:"id"`
	CommandType CommandType `json:"commandType"`
	Time        int64       `json:"time"`
	SendTime    int64       `json:"sendTime"`
	Type        EventType   `json:"type"`
	MallID      int64       `json:"mallID"`
}

func (ack AckMessage) sendAck(conn net.Conn) error {
	data, err := json.Marshal(&ack)
	if err != nil {
		return err
	}
	return wsutil.WriteClientText(conn, data)
}

func NewAckMessage(id uint64, sendTime int64, eventType EventType, mallID int64) *AckMessage {
	return &AckMessage{
		ID:          id,
		CommandType: CommandType_ACK,
		Time:        time.Now().UnixMilli(),
		SendTime:    sendTime,
		Type:        eventType,
		MallID:      mallID,
	}
}

type HeartBeatMessage struct {
	CommandType CommandType `json:"commandType"`
	Time        int64       `json:"time"`
}

func (hb HeartBeatMessage) sendHeartBeat(conn net.Conn) error {
	data, err := json.Marshal(&hb)
	if err != nil {
		return err
	}
	return wsutil.WriteClientText(conn, data)
}

func NewHeartBeatMessage() *HeartBeatMessage {
	return &HeartBeatMessage{
		CommandType: CommandType_HEARTBEAT,
		Time:        time.Now().UnixMilli(),
	}
}

type HandleResult chan error
