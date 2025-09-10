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

func (s *StudentRepository) GetStudentData(reqContext context.Context, offset, limit int) ([]models.StudentData, error) {
	sql := "SELECT id, name, role FROM students LIMIT $2 OFFSET $1"
	values := []any{offset, limit}
	rows, err := s.db.Query(reqContext, sql, values...)
	if err != nil {
		log.Println("Internal Server Error: ", err.Error())
		return []models.StudentData{}, err
	}
	defer rows.Close()
	var students []models.StudentData
	// membaca rows/record
	for rows.Next() {
		var student models.StudentData
		if err := rows.Scan(&student.Id, &student.Name, &student.Role); err != nil {
			log.Println("Scan Error, ", err.Error())
			return []models.StudentData{}, err
		}
		students = append(students, student)
	}
	return students, nil
}

// func (s *StudentRepository) Add(){}
func (s *StudentRepository) EditImage(rctx context.Context, images string, id int) (models.StudentData, error) {
	sql := "UPDATE students SET image=$1 WHERE id=$2 RETURNING id, name, image"
	values := []any{images, id}

	var student models.StudentData
	err := s.db.QueryRow(rctx, sql, values...).Scan(&student.Id, &student.Name, &student.Image)
	if err != nil {
		log.Println("Internal server error.\nCause: ", err.Error())
		return models.StudentData{}, err
	}
	return student, nil
}

func (s *StudentRepository) GetStudentById(rctx context.Context, id int) (models.StudentData, error) {
	sql := "SELECT id, name, image FROM students WHERE id = $1"
	values := []any{id}
	var student models.StudentData
	if err := s.db.QueryRow(rctx, sql, values...).Scan(&student.Id, &student.Name, &student.Image); err != nil {
		return models.StudentData{}, err
	}
	return student, nil
}
