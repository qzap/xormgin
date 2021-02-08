package main

import (
	"fmt"
	"github.com/qzap/xormgin/routing"
)

func main() {
	server := routing.WebService{}
	server.Run()
	fmt.Println("Server started!")
}
