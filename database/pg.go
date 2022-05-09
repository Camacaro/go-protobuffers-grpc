package database

import (
	"context"
	"database/sql"
	"go-grpc/models"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing rows: ", err)
		}
	}()

	var student = models.Student{}

	for rows.Next() {
		err := rows.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}

		return &student, nil
	}

	return &student, nil
}

func (r *PostgresRepository) SetStudent(ctx context.Context, student *models.Student) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO students (id, name, age) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	return err
}

func (r *PostgresRepository) SetTest(ctx context.Context, test *models.Test) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO tests (id, name) VALUES ($1, $2)", test.Id, test.Name)
	return err
}

func (r *PostgresRepository) GetTest(ctx context.Context, id string) (*models.Test, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM tests WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing rows: ", err)
		}
	}()

	var test = models.Test{}

	for rows.Next() {
		err := rows.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}

		return &test, nil
	}

	return &test, nil
}

func (r *PostgresRepository) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO questions (id, answer, question, test_id) VALUES ($1, $2, $3, $4)", question.Id, question.Answer, question.Question, question.TestId)
	return err
}

func (r *PostgresRepository) GetQuestion(ctx context.Context, id string) (*models.Question, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, answer, question, test_id FROM questions WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing rows: ", err)
		}
	}()

	var question = models.Question{}

	for rows.Next() {
		err := rows.Scan(&question.Id, &question.Answer, &question.Question, &question.TestId)
		if err != nil {
			return nil, err
		}

		return &question, nil
	}

	return &question, nil
}

func (repo *PostgresRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	// rows, err := repo.db.QueryContext(ctx, "SELECT students.id, students.name, students.age FROM students INNER JOIN enrollments ON students.id = enrollments.student_id WHERE enrollments.test_id = $1", testId)
	// Estudiantes dentro del test
	rows, err := repo.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)", testId)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal("Error closing rows: ", err)
		}
	}()

	var students []*models.Student

	for rows.Next() {
		var student = models.Student{}
		if err = rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
