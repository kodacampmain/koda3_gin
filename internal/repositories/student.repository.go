package repositories

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/models"
)

type StudentRepository struct {
	db *pgxpool.Pool
}

func NewStudentRepository(db *pgxpool.Pool) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (s *StudentRepository) GetStudentData(reqContext context.Context, offset, limit int) ([]models.Student, error) {
	sql := "SELECT id, name, role FROM students LIMIT $2 OFFSET $1"
	values := []any{offset, limit}
	rows, err := s.db.Query(reqContext, sql, values...)
	if err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return []models.Student{}, err
	}
	defer rows.Close()
	var students []models.Student
	// membaca rows/record
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.Id, &student.Name, &student.Role); err != nil {
			log.Println("Scan Error, ", err.Error())
			return []models.Student{}, err
		}
		students = append(students, student)
	}
	return students, nil
}

// func (s *StudentRepository) Add(){}
