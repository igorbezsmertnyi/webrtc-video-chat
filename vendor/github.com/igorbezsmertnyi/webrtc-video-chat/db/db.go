package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS rooms (
      id       SERIAL,
			slug		 TEXT,
			peer		 TEXT,
      created_at TIMESTAMP
    );
  `)

	if err != nil {
		return nil, err
	}

	return db, nil
}
