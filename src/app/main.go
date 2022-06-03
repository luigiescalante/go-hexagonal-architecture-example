package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.hex-architecture/internal/core/service"
	"go.hex-architecture/internal/infrastructure/http/gin_engine"
	"go.hex-architecture/internal/infrastructure/repository/cmd"
	"go.hex-architecture/internal/infrastructure/repository/file"
	"go.hex-architecture/internal/infrastructure/repository/mysql"
	"log"
	"os"
	"strings"
)

func main() {
	err := godotenv.Load("./app/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	cmd.ClearScreen()
	fmt.Println("Register a list of names:")
	fmt.Println("Choose an option to save the list")
	fmt.Println("Select an option")
	fmt.Println("1) Console")
	fmt.Println("2) File")
	fmt.Println("3) Database")
	fmt.Println("4) Webservice")
	var option int
	_, err = fmt.Scanln(&option)
	if err != nil {
		log.Println(err)
		return
	}
	var helloMsg *service.HelloMessageSrv
	switch option {
	case 1:
		helloMsg = service.NewHelloMessage(cmd.NewHelloMessageRepo())
	case 2:
		helloMsg = service.NewHelloMessage(file.NewHelloMessageRepo())
	case 3:
		helloMsg = service.NewHelloMessage(mysql.NewHelloMessageRepo())
	case 4:
		log.Println("Use a request tool for test webservice")
		go func() {
			var server *service.HttpServerSrv
			_, err = server.Server(gin_engine.Server())
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}()
	default:
		fmt.Println("option not valid.")
		os.Exit(0)
	}
	condition := true
	var name string
	var controlList string
	for ok := true; ok; ok = condition {
		fmt.Println("Write a name:")
		_, err = fmt.Scanln(&name)
		if err != nil {
			log.Println(err)
			return
		}
		helloMsg.CreateMessage(name)
		err = helloMsg.Repository.Save(helloMsg.HelloMsg.Message)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Another name: s/n")
		_, err = fmt.Scanln(&controlList)
		if err != nil {
			log.Println(err)
			return
		}
		if strings.ToLower(controlList) != "s" {
			condition = false
		}
		cmd.ClearScreen()
	}
	list := helloMsg.Repository.GetList()
	for _, message := range list {
		fmt.Println(message)
	}
}
