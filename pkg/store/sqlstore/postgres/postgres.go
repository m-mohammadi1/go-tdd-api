package postgres

import "database/sql"

type postgresSore struct {
	db *sql.DB
}

func NewPostgresStore(db *sql.DB) *postgresSore {
	return &postgresSore{
		db: db,
	}
}
