package repository

import (
	"database/sql"
	"schoolApplication/models"
)

type StudentRepository interface {
	Create(student *models.Student) error
	GetAll() ([]models.Student, error)
	Update(student *models.Student) error
	Delete(id int) error
}

type PostgresStudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) StudentRepository {
	return &PostgresStudentRepository{DB: db}
}

func (r *PostgresStudentRepository) Create(student *models.Student) error {
	sqlStatement := `INSERT INTO students (Name, Surname, Email) VALUES ($1, $2, $3) RETURNING ID`
	err := r.DB.QueryRow(sqlStatement, student.Name, student.Surname, student.Email).Scan(&student.ID)
	return err
}

func (r *PostgresStudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query(`SELECT ID, Name, Surname, Email FROM students`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := make([]models.Student, 0)
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.Name, &s.Surname, &s.Email); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

func (r *PostgresStudentRepository) Update(student *models.Student) error {
	sqlStatement := `UPDATE students SET Name=$1, Surname=$2, Email=$3 WHERE ID=$4`
	_, err := r.DB.Exec(sqlStatement, student.Name, student.Surname, student.Email, student.ID)
	return err
}

func (r *PostgresStudentRepository) Delete(id int) error {
	sqlStatement := `DELETE FROM students WHERE ID=$1`
	_, err := r.DB.Exec(sqlStatement, id)
	return err
}
