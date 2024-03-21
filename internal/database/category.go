package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("Insert INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("Select id, name, description From categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Name: name, Description: description})

	}

	return categories, nil
}

func (c *Category) FindByCourseID(id string) (Category, error) {
	rows, err := c.db.Query("Select ct.id, ct.name, ct.description From categories ct inner join courses co On ct.id = co.category_id where co.id = $1", id)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return Category{}, err
		}
		return Category{ID: id, Name: name, Description: description}, nil

	}

	return Category{}, nil
}

func (c *Category) FindByID(id string) (Category, error) {
	rows, err := c.db.Query("Select id, name, description From categories where id = $1", id)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			return Category{}, err
		}
		return Category{ID: id, Name: name, Description: description}, nil

	}

	return Category{}, nil
}
