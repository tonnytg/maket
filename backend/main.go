package main

import (
	"github.com/tonnytg/makemoneytarget/internal/infra/database"
	"github.com/tonnytg/makemoneytarget/internal/infra/webserver"
)

func main() {

	database.Start()

	webserver.Start()
}
