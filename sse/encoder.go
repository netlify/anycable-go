package sse

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/anycable/anycable-go/common"
	"github.com/anycable/anycable-go/encoders"
	"github.com/anycable/anycable-go/ws"
)

const sseEncoderID = "sse"

// Tell the client to reconnect in a year in case we don't really want it to re-connect
const retryNoReconnect = 31536000000

// Encoder is responsible for converting messages to SSE format (event:, data:, etc.)
// NOTE: It's only used to encode messages from server to client.
type Encoder struct {
}

func (Encoder) ID() string {
	return sseEncoderID
}

func (Encoder) Encode(msg encoders.EncodedMessage) (*ws.SentFrame, error) {
	msgType := msg.GetType()

	b, err := json.Marshal(&msg)
	if err != nil {
		panic("Failed to build JSON 😲")
	}

	payload := "data: " + string(b)
	if msgType != "" {
		payload = "event: " + msgType + "\n" + payload
	}

	if reply, ok := msg.(*common.Reply); ok {
		if reply.Offset > 0 && reply.Epoch != "" && reply.StreamID != "" {
			payload += "\nid: " + fmt.Sprintf("%d/%s/%s", reply.Offset, reply.Epoch, reply.StreamID)
		}
	}

	if msgType == "disconnect" {
		dmsg, ok := msg.(*common.DisconnectMessage)
		if ok && !dmsg.Reconnect {
			payload += "\nretry: " + fmt.Sprintf("%d", retryNoReconnect)
		}
	}

	return &ws.SentFrame{FrameType: ws.TextFrame, Payload: []byte(payload)}, nil
}

func (e Encoder) EncodeTransmission(raw string) (*ws.SentFrame, error) {
	msg := common.Reply{}

	if err := json.Unmarshal([]byte(raw), &msg); err != nil {
		return nil, err
	}

	return e.Encode(&msg)
}

func (Encoder) Decode(raw []byte) (*common.Message, error) {
	return nil, errors.New("unsupported")
}
