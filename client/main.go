package main

import (
	"context"
	"go-grpc/testproto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Conectamos al server de test
	// Esta es una conexion insegura ya que nuestro servidor no tiene certificado, instalado ese Cert
	// en caso de conectarme a un servidor encriotado la opcion que tengo que modificar es el segundo parametro
	// del Dial
	cc, err := grpc.Dial("localhost:5070", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	// Cerrar la coneccion cuando termine el main
	defer cc.Close()

	c := testproto.NewTestServiceClient(cc)

	// DoUnary(c)
	// DoClientStreaming(c)
	// DoServerStreaming(c)
	DoBidireccionalStreaming(c)
}

// Funcion cliente unario, esto es lo mas parecito a REST api
func DoUnary(c testproto.TestServiceClient) {
	req := &testproto.GetTestRequest{
		Id: "t1",
	}

	// La invocacion de hace aqui, pero esta implementacion se hace en el server
	res, err := c.GetTest(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not get test: %v", err)
	}

	log.Printf("Test Answer: %v", res)
}

func DoClientStreaming(c testproto.TestServiceClient) {
	questions := []*testproto.Question{
		{
			Id:       "q1",
			Answer:   "Azul",
			Question: "Colo asociado a golang",
			TestId:   "t1",
		},
		{
			Id:       "q2",
			Answer:   "Google",
			Question: "Empresa que creo golang",
			TestId:   "t1",
		},
		{
			Id:       "q3",
			Answer:   "Backend",
			Question: "Espacialidad Golang",
			TestId:   "t1",
		},
	}

	stream, err := c.SetQuestion(context.Background())

	if err != nil {
		log.Fatalf("Could not set question: %v", err)
	}

	for _, q := range questions {
		log.Println("Sending question: ", q.Question)

		// Enviar las preguntas
		stream.Send(q)

		// Simular un retraso
		time.Sleep(2 * time.Second)
	}

	// He finalizado el envio de mensajes y espero la respuesta de el server
	msg, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Could not CloseAndRecv: %v", err)
	}

	log.Printf("Server Answer: %v", msg)
}

func DoServerStreaming(c testproto.TestServiceClient) {
	req := &testproto.GetStudentsPerTestRequest{
		TestId: "t1",
	}

	stream, err := c.GetStudentsPerTest(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GetStudentsPerTest: %v", err)
	}

	// Indefinida porque no sabemos cuantos estudiantes vamos a recibi
	for {
		student, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}

		log.Printf("Student: %v", student)
	}
}

func DoBidireccionalStreaming(c testproto.TestServiceClient) {
	answer := testproto.TakeTestRequest{
		Answer: "Azul",
	}

	numberOfQuestions := 3

	// Para emular el comportamiento de una persona que contesta una pregunta
	// y el server responde con una pregunta

	/*
		Para emular el comportamiento de enviar un streaming y
		recibir un streaming vamos a crear una goRoutina para
		cada escenario
	*/

	// Nos va a servir para un chanel de control para bloquear el programa al ejecutar las goroutinas
	waitChannel := make(chan struct{})

	stream, err := c.TakeTest(context.Background())
	if err != nil {
		log.Fatalf("error while calling TakeTest: %v", err)
	}

	// Crear una goroutine para enviar las respuesta
	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			log.Printf("Sending answer: %v", answer.Answer)
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()

	// Crear una goroutine para recibir las preguntas
	go func() {
		// Indefinidas porque no sabemos cuantas veces va a responder el servidor
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while reading stream: %v", err)
				break
			}

			log.Printf("Responde from server: %v", res.Question)
		}

		close(waitChannel)
	}()

	// Esperar que las goroutinas terminen
	<-waitChannel
}
