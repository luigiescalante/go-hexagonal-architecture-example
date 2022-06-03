package service

import "go.hex-architecture/internal/core/ports"

type HttpServerSrv struct {
	engine interface{}
}

func (server HttpServerSrv) Server(engine ports.ServerHttpSrv) (*HttpServerSrv, error) {
	return &HttpServerSrv{
		engine: engine,
	}, nil
}

func (server HttpServerSrv) Start() error {
	return nil
}
