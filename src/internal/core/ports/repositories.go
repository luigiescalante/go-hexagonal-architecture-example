package ports

type HelloMessageRepo interface {
	Save(message string) error
	GetMessage() string
}
