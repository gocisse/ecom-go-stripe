package models

import (
	"context"
	"database/sql"
	"time"
)

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
//function to get one widget item from the database
func (m *DBModel) GetWidget(id int) (Widget, error) {
	var widget Widget

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, "select id, name from widgets where id = ?", id)
	err := row.Scan(&widget.ID, &widget.Name)
	if err != nil {
		return widget, err
	}

	return widget, err
}
