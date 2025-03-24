package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDatabase() {
	var err error
	DB, err = sqlx.Connect("postgres", "host=localhost user=postgres password=postgres dbname=db_go sslmode=disable")
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Database connected!")
}
