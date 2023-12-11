package main

import (
	"log"

	"cardap.in/lambda/db"
)

func main() {
	log.Println("Cardapin WebServer")

	log.Println("Initializing database connection")
	c := db.PostgresConnector{}
	c.InitializeDatabase()

	log.Println("Initializing application")
	a := App{}
	a.Initialize()

	log.Println("Running webserver")
	a.Run(":8080")
}
