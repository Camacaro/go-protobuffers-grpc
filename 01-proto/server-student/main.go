package main

import (
	"go-grpc/database"
	"go-grpc/server"
	"go-grpc/studentproto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Vamos a crear un listener de una coneccion TCP
	list, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	// Creamos un nuevo server gRPC
	server := server.NewStudentServer(repo) // Server que implementa nuesto gRPC
	s := grpc.NewServer()                   // Connecion gRPC
	studentproto.RegisterStudentServiceServer(s, server)

	/*
		Nos va a servir para decirle a los clientes (postman)
		que podemos proveer cierta metadata para poder testear
		la API y va a deoender de lo que reflection le este dando de data
	*/
	reflection.Register(s) // Registramos el server para que sepa que estamos usando el protobuffer

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
