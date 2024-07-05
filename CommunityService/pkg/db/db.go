package db

import (
	"database/sql"
	"fmt"
	"log"
)

func DB() (*sql.DB, error) {

	var psqlUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"dodi",
		"ecotruck",
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
		return nil, err
	}

	return psqlConn, nil
}