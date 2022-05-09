package server

import (
	"context"
	"go-grpc/models"
	"go-grpc/repository"
	"go-grpc/testproto"
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

// func (s *TestServer) SetQuestions(stream testproto.TestService_SetQuestionsServer) error {
// 	for {
// 		msg, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&testproto.SetQuestionResponse{
// 				Ok: true,
// 			})
// 		}
// 		if err != nil {
// 			log.Fatalf("Error reading stream: %v", err)
// 			return err
// 		}
// 		question := &models.Question{
// 			Id:       msg.GetId(),
// 			Answer:   msg.GetAnswer(),
// 			Question: msg.GetQuestion(),
// 			TestId:   msg.GetTestId(),
// 		}
// 		err = s.repo.SetQuestion(context.Background(), question)
// 		if err != nil {
// 			return stream.SendAndClose(&testproto.SetQuestionResponse{
// 				Ok: false,
// 			})
// 		}

// 	}
// }

// func (s *TestServer) EnrollStudents(stream testproto.TestService_EnrollStudentsServer) error {
// 	for {
// 		msg, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&testproto.SetQuestionResponse{
// 				Ok: true,
// 			})
// 		}
// 		if err != nil {
// 			log.Fatalf("Error reading stream: %v", err)
// 			return err
// 		}
// 		enrollment := &models.Enrollment{
// 			StudentId: msg.GetStudentId(),
// 			TestId:    msg.GetTestId(),
// 		}
// 		err = s.repo.SetEnrollment(context.Background(), enrollment)
// 		if err != nil {
// 			return stream.SendAndClose(&testproto.SetQuestionResponse{
// 				Ok: false,
// 			})
// 		}

// 	}
// }

// func (s *TestServer) GetStudentsPerTest(req *testproto.GetStudentsPerTestRequest, stream testproto.TestService_GetStudentsPerTestServer) error {
// 	students, err := s.repo.GetStudentsPerTest(context.Background(), req.GetTestId())
// 	if err != nil {
// 		return err
// 	}
// 	for _, student := range students {
// 		student := &studentpb.Student{
// 			Id:   student.Id,
// 			Name: student.Name,
// 			Age:  student.Age,
// 		}
// 		err := stream.Send(student)
// 		time.Sleep(2 * time.Second)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (s *TestServer) TakeTest(stream testproto.TestService_TakeTestServer) error {
// 	questions, err := s.repo.GetQuestionsPerTest(context.Background(), "t1")
// 	if err != nil {
// 		return err
// 	}
// 	i := 0
// 	var currentQuestion = &models.Question{}
// 	for {
// 		if i < len(questions) {
// 			currentQuestion = questions[i]
// 		}

// 		if i <= len(questions)-1 {
// 			questionToSend := &testproto.Question{
// 				Id:       currentQuestion.Id,
// 				Question: currentQuestion.Question,
// 			}
// 			err := stream.Send(questionToSend)
// 			if err != nil {
// 				log.Printf("Error sending question: %v", err)
// 				return err
// 			}
// 			i++
// 		}
// 		answer, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			log.Printf("Error receiving answer: %v", err)
// 			return err
// 		}
// 		log.Println("Answer for question:", currentQuestion.Question, "is", answer.GetAnswer())
// 	}
// }
