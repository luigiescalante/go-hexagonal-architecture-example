package ports

type HelloMessageRepo interface {
	Save(message string) error
	GetList() []string
}
