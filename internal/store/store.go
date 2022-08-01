package store

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type SourceRecord struct {
	ID         string `db:"id"`
	SourceName string `db:"source_name"`
	SourceURL  string `db:"source_url"`
	Category   string `db:"category"`
}

type Store struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *Store {
	return &Store{DB: db}
}

//TODO: use sqlc for generating queries/models

func (s *Store) GetAllSources() ([]SourceRecord, error) {
	query := `SELECT id, source_name, source_url, category FROM sources;`

	rows, err := s.DB.Queryx(query)

	if err != nil {
		log.Println("query failed: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	res := make([]SourceRecord, 0)
	var r SourceRecord
	for rows.Next() {
		err = rows.StructScan(&r)
		if err != nil {
			log.Println("could not read row: " + err.Error())
		}
		res = append(res, r)
	}
	return res, nil
}

//TODO: `AddSource` function
//TODO: integration tests.
