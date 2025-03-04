package main

import (
	"database/sql"
	"fmt"
	"github.com/hako/branca"
	"log"
)

const (
	databaseURL = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	port = 3000
)

func main() {
	fmt.Println("Tweety")

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Error opening db: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging db: %v\n", err)
		return
	}

	codec := branca.NewBranca("supersecretkeyyoushouldnotshare")
	s := service.Service {
		db: db,
		codec: codec,
	}
}