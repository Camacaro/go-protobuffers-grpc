package server

import (
	"context"
	"fmt"
	"go-grpc/models"
	"go-grpc/repository"
	"go-grpc/studentproto"
	"go-grpc/testproto"
	"io"
	"log"
	"time"
)

type TestServer struct {
	repo repository.Repository
	testproto.UnimplementedTestServiceServer
}

func NewTestServer(repo repository.Repository) *TestServer {
	return &TestServer{
		repo: repo,
	}
}

func (s *TestServer) GetTest(ctx context.Context, req *testproto.GetTestRequest) (*testproto.Test, error) {
	test, err := s.repo.GetTest(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &testproto.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testproto.Test) (*testproto.SetTestResponse, error) {
	test := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testproto.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

/*
	Esto es parte de la implemantacion de streaming del lado de cliente

	Cuando se abre la conexion se puede enviar todos los mensajes que quiera ,
	hasta que lo cierro y ahi el server me resonde
*/
func (s *TestServer) SetQuestion(stream testproto.TestService_SetQuestionServer) error {
	// Iteramos porque no sabemos con exactitud cuantas peticiones vamos a recibir, preguntas a recibir
	for {
		// Llamamos el stream, esto se va a bloquear hasta recibir una peticion del cliente
		msg, err := stream.Recv()

		// EOF: End of File, es una condicion de salida del cliente, el lo ha terminado y quiere la respuesta
		if err == io.EOF {
			// Esto es para terminar el stream y devolver la respuesta
			return stream.SendAndClose(&testproto.SetQuestionResponse{
				Ok: true,
			})
		}

		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}

		question := &models.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}

		err = s.repo.SetQuestion(context.Background(), question)

		if err != nil {
			log.Fatalf("Error setting question: %v", err)
			return stream.SendAndClose(&testproto.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

// Esto es parte de la implemantacion de streaming del lado de server
func (s *TestServer) EnrollStudents(stream testproto.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testproto.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}
		enrollment := &models.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}
		err = s.repo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testproto.SetQuestionResponse{
				Ok: false,
			})
		}

	}
}

/*
	Esto es parte de la implemantacion de streaming del lado de server
	Aqui tenemos que devolver un stream de datos constantemente

	El segundo parametro stream, es el que usaremos para devolver los datos
*/
func (s *TestServer) GetStudentsPerTest(req *testproto.GetStudentsPerTestRequest, stream testproto.TestService_GetStudentsPerTestServer) error {
	// Leeremos los estudiantes que estan inscritos en el test
	fmt.Println("GetStudentsPerTest", req.GetTestId())
	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, student := range students {
		student := &studentproto.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		// Enviamos el estudiante al cliente
		err := stream.Send(student)

		// Solo por practica, dar una simulacion de el server esta tardando
		time.Sleep(2 * time.Second)

		if err != nil {
			return err
		}
	}
	return nil
}

/*
	Esto es de la implementacion de streaming bidireccional

	stream, sirve para enviar y recibir datos,
	Tiene la interfaz de enviar(Send) y recibir(Recv)
*/
func (s *TestServer) TakeTest(stream testproto.TestService_TakeTestServer) error {
	// Aqui el id del test esta en duro
	questions, err := s.repo.GetQuestionsPerTest(context.Background(), "t1")

	if err != nil {
		return err
	}

	i := 0
	var currentQuestion = &models.Question{}
	for {
		if i < len(questions) {
			currentQuestion = questions[i]
		}

		if i <= len(questions)-1 {
			// Esta es la respuesta que le enviamos al cliente
			questionToSend := &testproto.Question{
				Id:       currentQuestion.Id,
				Question: currentQuestion.Question,
			}
			// Enviamos la pregunta al cliente
			err := stream.Send(questionToSend)

			if err != nil {
				log.Printf("Error sending question: %v", err)
				return err
			}
			i++
		}

		// Esperamos la respuesta del cliente
		answer, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving answer: %v", err)
			return err
		}
		log.Println("Answer for question:", currentQuestion.Question, "is", answer.GetAnswer())
	}
}
