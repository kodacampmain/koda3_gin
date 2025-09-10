package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kodacampmain/koda3_gin/internal/models"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) GetStudentWithPasswordAndRole(rctx context.Context, name string) (models.StudentData, error) {
	// validasi user
	// ambil data user berdasarkan input user
	sql := "SELECT id, name, password, role FROM students WHERE name = $1"

	var student models.StudentData
	if err := a.db.QueryRow(rctx, sql, name).Scan(&student.Id, &student.Name, &student.Password, &student.Role); err != nil {
		if err == pgx.ErrNoRows {
			return models.StudentData{}, errors.New("user not found")
		}
		log.Println("Internal Server Error.\nCause: ", err.Error())
		return models.StudentData{}, err
	}
	return student, nil
}
