package main

import (
	"service-backend/database"
	"service-backend/routes"
	"service-backend/shared"
)

func main() {

	shared.LogCustom([]string{"Iniciando o servidor goAPI"}, "info")

	database.ConnectDB()

	routes.HandleRequest()

}
