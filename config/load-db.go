package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type DbConfigType struct {
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
	dbPort     int
}

func DbConfig() (*DbConfigType, error) {

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dbPortNumber, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, fmt.Errorf("error al convertir DB_PORT a entero: %w", err)
	}

	return &DbConfigType{
		dbHost:     dbHost,
		dbName:     dbName,
		dbUser:     dbUser,
		dbPassword: dbPassword,
		dbPort:     dbPortNumber,
	}, nil

}

func DbLoad() (*sql.DB, error) {

	dbConfig, err := DbConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot get dbConfig func %w", err))
		return nil, err
	}
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbConfig.dbUser, dbConfig.dbPassword, dbConfig.dbName, dbConfig.dbHost, dbConfig.dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		defer db.Close()
		log.Fatal(fmt.Errorf("unable to connect to database: %w", err))
	}

	log.Println("Pinging database to check connection...")
	if err := db.Ping(); err != nil {
		defer db.Close()
		log.Fatal(fmt.Errorf("ping to database failed: %w", err))
		return nil, err
	}
	log.Println("Successfully connected and pinged the database!")
	return db, nil
}
