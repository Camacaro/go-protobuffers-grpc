package server

import (
	"context"
	"go-grpc/models"
	"go-grpc/repository"
	"go-grpc/studentproto"
)

type Server struct {
	repo                                           repository.Repository
	studentproto.UnimplementedStudentServiceServer // Composicion sobre la herencia, esto es necesario paor poder implementar la interfaz en el server-student
}

func NewStudentServer(repo repository.Repository) *Server {
	return &Server{repo: repo}
}

// Aqui hacemos referencia al archivo studentproto.go (protobuffer)
func (s *Server) GetStudent(ctx context.Context, req *studentproto.GetStudentRequest) (*studentproto.Student, error) {
	// req.Id es el id que se envio en el request del protobuffer (GetStudentRequest)
	// puedo usar req.Id o req.GetId()
	student, err := s.repo.GetStudent(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// devolvemos el Student, que necesitamos en la interfaz del studentproto.Student
	return &studentproto.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentproto.Student) (*studentproto.SetStudentResponse, error) {
	student := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return &studentproto.SetStudentResponse{
		Id: student.Id,
	}, nil
}
