package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	dbUsername := os.Getenv("RDS_DB_USERNAME")
	dbPassword := os.Getenv("RDS_DB_PASSWORD")
	dbHost := os.Getenv("RDS_DB_HOST")
	dbName := os.Getenv("RDS_DB_NAME")
	dbPort := os.Getenv("RDS_DB_PORT")

	return &Config{
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbName:     dbName,
		DbPort:     dbPort,
	}
}

func MySQL(c *Config) (*sql.DB, error) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
	db, err := sql.Open("mysql", dbUri)

	if err != nil {
		return nil, err
	}

	log.Println("Successfully to connect in database")
	return db, nil
}
