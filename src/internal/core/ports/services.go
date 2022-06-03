package ports

type HelloMessageSrv interface {
	CreateMessage(msg string) string
}

type ServerHttpSrv interface {
}
