package ping

import (
	"context"
	sprut_server "sprut/gen/server"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) GetPing(ctx context.Context) (sprut_server.GetPingRes, error) {
	return &sprut_server.Pong{
		Pong: true,
	}, nil
}
