package file

import (
	"bufio"
	"log"
	"os"
)

const filePath = "./app/list-name.csv"

type repository struct {
	file *os.File
}

func NewHelloMessageRepo() *repository {
	return &repository{}
}

func (r *repository) Save(message string) error {
	err := r.CreateFile()
	if err != nil {
		return err
	}
	_, err = r.file.WriteString(message + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetList() []string {
	var list []string
	var err error
	r.file, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(r.file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	defer r.file.Close()
	return list
}

func (r *repository) CreateFile() error {
	var err error
	r.file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	return nil

}
