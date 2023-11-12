package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ps := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		`localhost`, `notes`, `notes_pass`, `notes`)

	db, err := sql.Open("pgx", ps)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
