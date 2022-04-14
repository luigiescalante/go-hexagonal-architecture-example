package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.hex-architecture/internal/core/service"
	"go.hex-architecture/internal/infrastructure/repository/cmd"
	"go.hex-architecture/internal/infrastructure/repository/file"
	"log"
)

func main() {
	err := godotenv.Load("./app/.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var name string
	var option int
	fmt.Println("Hello write your name:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Select an option")
	fmt.Println("1) Console")
	fmt.Println("2) File")
	fmt.Println("3) WebService")
	_, err = fmt.Scanln(&option)
	if err != nil {
		return
	}
	var helloMsg *service.HelloMessageSrv
	switch option {
	case 1:
		helloMsg = service.NewHelloMessage(cmd.NewHelloMessageRepo())
	case 2:
		helloMsg = service.NewHelloMessage(file.NewHelloMessageRepo())
	case 3:
		log.Println("pendiente")
	default:
		fmt.Println("option not valid.")
	}
	helloMsg.CreateMessage(name)
	helloMsg.Repository.Save(helloMsg.HelloMsg.Message)
	message := helloMsg.GetMessage()
	log.Println(message)
	/*router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		helloSrv.CreateMessage("kaiba")
		err := helloSrv.Repository.Save(helloSrv.HelloMsg.Message)
		if err != nil {
			return
		}
		helloSrv = service.NewHelloMessage(cmd.NewHelloMessageRepo())
		helloSrv.CreateMessage("zero ")
		err = helloSrv.Repository.Save(helloSrv.HelloMsg.Message)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"system": helloSrv.HelloMsg.Message,
		})
	})
	router.Run()*/
}
