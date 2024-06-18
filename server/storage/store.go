package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	Connection *string
	DB *sql.DB

}

func New(connection string) *Store {
	store := Store{
		Connection: &connection,
	}
	return &store
}

func (s *Store) Open() error {

	db, err := sql.Open("postgres", *s.Connection)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil{
		return err
	}

	
	s.DB = db

	return nil
}