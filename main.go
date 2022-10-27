package main

import (
	"github.com/LucasGabrielBecker/go-rest-api-gin/database"
	"github.com/LucasGabrielBecker/go-rest-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
