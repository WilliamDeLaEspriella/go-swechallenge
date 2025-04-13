package main

import (
	"github.com/WilliamDeLaEspriella/go-swechallenge/app"
)

func main() {
	var a app.Server
	a.CreateConnection()
	a.Migrate()
	a.CreateRoutes()
	a.Run()
}
