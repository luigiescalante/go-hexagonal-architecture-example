package gin_engine

import (
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
}

func Server() *HttpServer {
	ginGonic := gin.Default()
	ginGonic.Run()
	return &HttpServer{
		engine: ginGonic,
	}
}

func (server *HttpServer) Start() error {
	/*log.Println("kaiba si entre")
	err := server.engine.Run()
	if err != nil {
		return nil
	}*/
	return nil
}
