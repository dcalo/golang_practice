package main

import (
	"log"

	myPack1 "example.com/myProject/adapters"
)

func main() {
	server := myPack1.NewServer()
	if err := server.Run(":3000"); err != nil {

		log.Fatalln("Can't start the server")

	}

}
