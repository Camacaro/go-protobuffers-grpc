package repository

import (
	"context"
	"go-grpc/models"
)

type Repository interface {
	GetStudents(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

// Se usa en tiempo de ejecucion
var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func SetStudent(ctx context.Context, student *models.Student) error {
	return implementation.SetStudent(ctx, student)
}

func GetStudents(ctx context.Context, id string) (*models.Student, error) {
	return implementation.GetStudents(ctx, id)
}
