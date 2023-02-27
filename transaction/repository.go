package transaction

import db "transationAPI/database"

type Repository struct {
	DB db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{
		DB: db,
	}
}
