package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type PostgresConnector struct {
}

type Connector struct {
}

func (p *PostgresConnector) InitializeDatabase() {
	var err error
	DB, err = p.getConnection()
	DB.Logger.LogMode(logger.Info)
	if err != nil {
		log.Printf("Error to get database connection: " + err.Error())
	}
}

func (c *Connector) InitializeDatabaseParam(db *gorm.DB) {
	DB = db
}

func (p *PostgresConnector) getConnection() (db *gorm.DB, err error) {
	godotenv.Load()
	e := godotenv.Load()
	if e != nil {
		oe := godotenv.Load("bin/.env")
		if oe != nil {
			log.Println(e)
		}
	}
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	if dbPort == "" {
		dbPort = "5432"
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	db2, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db2, err
}
