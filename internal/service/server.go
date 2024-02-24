package service

import (
	"context"
	sprut_server "sprut/gen/server"
	"sprut/internal/service/routes/ping"
)

type serverImpl struct {
	*ping.PingHandler
}

func NewServerImpl() *sprut_server.Server {
	impl := &serverImpl{PingHandler: ping.NewPingHandler()}
	server, _ := sprut_server.NewServer(impl)
	return server
}

func (h *serverImpl) NewError(ctx context.Context, err error) *sprut_server.ErrorResponseStatusCode {
	return &sprut_server.ErrorResponseStatusCode{
		StatusCode: 500,
		Response: sprut_server.Error{
			Code:    500,
			Message: "Unexpected error!",
		},
	}
}


