package models

import "database/sql"

// DBModels is the type for a database connections values

type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for all models
type Models struct {
	DB DBModel
}

// Function to handle database connection
// NewModels return a model tupe with a database connection pool
func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

// Create type to hold a widget to store in the database

type Widget struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	InventoryLevel int    `json:"inventory_level"`
	Price          int    `json:"price"`
	CeatedAt       int    `json:"-"`
	UpdatedAt      int    `json:"-"`
}
