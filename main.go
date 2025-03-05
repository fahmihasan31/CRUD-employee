package main

import (
	"github/fahmihasan31/crud-employee-go/database"
	"github/fahmihasan31/crud-employee-go/routes"
	"net/http"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
