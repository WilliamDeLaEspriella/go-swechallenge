package main

import (
	"github.com/WilliamDeLaEspriella/go-swechallenge/app"
)

func main() {
	var server app.Server
	server.CreateConnection()
	server.CreateTables()
	server.Migrate()
	server.CreateRoutes()
	server.Run()
}
