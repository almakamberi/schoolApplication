package repository

import (
	"database/sql"
	"schoolApplication/models"
)

type ClassRepository interface {
	Create(class *models.Class) error
	GetAll() ([]models.Class, error)
	Update(class *models.Class) error
	Delete(id int) error
}

type PostgresClassRepository struct {
	DB *sql.DB
}

func NewClassRepository(db *sql.DB) ClassRepository {
	return &PostgresClassRepository{DB: db}
}

func (r *PostgresClassRepository) Create(class *models.Class) error {
	sqlStatement := `INSERT INTO class (name, date_of_creation) VALUES ($1, $2) RETURNING ID`
	err := r.DB.QueryRow(sqlStatement, class.Name, class.Date_Of_Creation).Scan(&class.ID)
	return err
}

func (r *PostgresClassRepository) GetAll() ([]models.Class, error) {
	rows, err := r.DB.Query(`SELECT ID, name, date_of_creation FROM class`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := make([]models.Class, 0)
	for rows.Next() {
		var c models.Class
		if err := rows.Scan(&c.ID, &c.Name, &c.Date_Of_Creation); err != nil {
			return nil, err
		}
		classes = append(classes, c)
	}

	return classes, nil
}

func (r *PostgresClassRepository) Update(class *models.Class) error {
	sqlStatement := `UPDATE class SET name=$1, date_of_creation=$2 WHERE ID=$3`
	_, err := r.DB.Exec(sqlStatement, class.Name, class.Date_Of_Creation, class.ID)
	return err
}

func (r *PostgresClassRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM class WHERE ID=$1`
	_, err := r.DB.Exec(sqlStatement, id)
	return err
}
