package transaction

import (
	db "transationAPI/database"
	"transationAPI/models"
)

type Repository struct {
	DB db.Database
}

func NewRepository(db db.Database) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetBalance() (*float64, error) {
	var transaction models.Transaction
	if err := r.DB.Order("id desc").First(&transaction).Error(); err != nil {
		return nil, err
	}
	return &transaction.Balance, nil
}

func (r *Repository) Create(trans *models.Transaction) error {
	// Start a new transaction
	tx := r.DB.Begin()

	// Lock the transactions table for writing
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Create(trans).Error(); err != nil {
		// Rollback the transaction if an error occurs
		tx.Rollback()
		return err
	}

	// Commit the transaction if there were no errors
	return tx.Commit().Error()
}
