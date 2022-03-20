package main

import (
	"log"
)

func main() {
	R.GET("/", index)
	R.POST("/create", create)

	log.Fatalln("Router encountered and error while main.Run:", R.Run(":"+Server_Port))
}

//
