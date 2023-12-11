package handlers

import (
	"log"

	"cardap.in/lambda/db"
)

func InitCtx() {
	c := db.PostgresConnector{}
	log.Println("Initializing database connection")
	c.InitializeDatabase()
}
