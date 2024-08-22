package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

type DatabaseStruct struct {
	Conn *sql.DB
}

func NewPostgresDB(config DatabaseConfig) (*DatabaseStruct, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Name, config.SSLMode)

	log.Printf("Attempting to connect to database with connection string: %s", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error opening database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to database")

	return &DatabaseStruct{Conn: db}, nil
}

func (db *DatabaseStruct) Close() error {
	return db.Conn.Close()
}
