package repository

import (
	"database/sql"
	"log"
	"material_storage/models"
)

type Material interface {
	Add(title, description, ref string) error
	List() ([]*models.Material, error)
}

type orders struct {
	db *sql.DB
}

func NewMaterialRepository(db *sql.DB) Material {
	return &orders{db: db}
}

func (r *orders) Add(title, description, ref string) error {
	_, err := r.db.Exec("INSERT INTO material (title, description, ref) VALUES ($1, $2, $3)", title, description, ref)
	if err != nil {
		return err
	}
	log.Println("Material created successfully")
	return nil
}

func (r *orders) List() ([]*models.Material, error) {
	rows, err := r.db.Query("SELECT title, description, ref, date_created FROM material")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	materials := make([]*models.Material, 0)
	for rows.Next() {
		m := new(models.Material)
		err := rows.Scan(&m.Title, &m.Description, &m.Ref, &m.DateCreated)
		if err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return materials, nil
}
