package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	dbStr := os.Getenv("DB_USER_PG_STRING")
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatalf("Can't connect to user database.\n[ERR]: %v", err)
	}

	// Connection pool
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(2 * time.Hour)

	// Ping to db
	if err := db.Ping(); err != nil {
		log.Fatalf("User DB not ready...\n[ERR]: %v", err)
	} else {
		log.Println("User DB connected successfully...")
	}

	return db
}
