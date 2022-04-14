package cmd

import "log"

type repository struct {
}

func NewHelloMessageRepo() *repository {
	return &repository{}
}

func (r *repository) Save(message string) error {
	log.Println("on file repo" + message)
	return nil
}

func (r *repository) GetMessage() string {
	log.Println("get message")
	return "kaiba"
}
