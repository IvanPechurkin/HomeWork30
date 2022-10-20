package main

import (
	"30/cmd/utils"
	"30/pkg/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Какой первый порт создаем")
	port1 := ""
	fmt.Scan(&port1)
	fmt.Println("Какой второй порт создаем")
	port2 := ""
	fmt.Scan(&port2)

	config1 := utils.Configuration{
		Database: utils.DatabaseSetting{
			Url:        "mongodb://localhost:27017",
			DbName:     "userdb",
			Collection: "user",
		},
		Server: utils.ServerSettings{
			Port: port1,
		},
	}

	config2 := utils.Configuration{
		Database: utils.DatabaseSetting{
			Url:        "mongodb://localhost:27017",
			DbName:     "userdb",
			Collection: "user",
		},
		Server: utils.ServerSettings{
			Port: port2,
		},
	}

	r1 := server.Init(config1)
	r2 := server.Init(config2)

	//Старт сервиса
	go func() {
		fmt.Printf("Started listening on port %s\n", port1)

		log.Fatal(http.ListenAndServe(port1, r1))
	}()
	go func() {
		fmt.Printf("Started listening on port %s\n", port2)

		log.Fatal(http.ListenAndServe(port2, r2))
	}()
	select {}
}
