package main

import (
	"go-grpc/database"
	"go-grpc/testproto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Vamos a crear un listener de una coneccion TCP
	list, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	// Creamos un nuevo server gRPC
	server := server.NewTestServer(repo) // Server que implementa nuesto gRPC
	s := grpc.NewServer()                // Connecion gRPC
	testproto.RegisterTestServiceServer(s, server)

	/*
		Nos va a servir para decirle a los clientes (postman)
		que podemos proveer cierta metadata para poder testear
		la API y va a deoender de lo que reflection le este dando de data

		Dentro del postman nos va a pedir los metodos
		"
			Load services and methods
			Choose how you want to load the services and methods.
			You can use a protobuf definition or load them using
			server reflection.
		"
		Por eso usamos reflection, al hacerlo nos va a deplegar las
		dos metodos que tenemos en el server (studentproto.go)
		y que estan en el package studentproto SetStudent y GetStudent

		PD: la respuesta es application/grpc
	*/
	reflection.Register(s) // Registramos el server para que sepa que estamos usando el protobuffer

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
