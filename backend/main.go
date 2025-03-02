package main

import (
	"github.com/tonnytg/makemoneytarget/internal/infra/webserver"
	"log"
)

func main() {
	log.Println("Start MakeMoneyTarget Backend")

	webserver.Start()
}
