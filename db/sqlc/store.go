package db

import "database/sql"

// SQLStore provides all functions to execute SQL queries
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creaetes a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
