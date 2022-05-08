# Protobuffers y gRPC
Nos ayuda a crear microservicios de manera más rápida y de alto rendimiento.

## Protobuffers
Es un protocolo de buffers (algo así como json), es un serializado de datos como XML, nos permite armar estructura de datos. 

message -> Nos permite definir un mensage -- Vamos a crear un mensaje o una structura nueva
Nombre {} -> dentro de los parentesis creamos la estructua de datos, tipo de la variable (string) y el nombre de la variable y se le asigna un valor que tiene que ser unico, ademas tiene que ser secuencial con las demas variables (estructua de dato), sera la posicion donde será serializado

message Student {
  string id = 1
  string name = 2  
}

Seralizar -> Deserializar 
Ya que como en el protobuffer se encuentran estructurada este proceso de deserialización será más rápida.

la extencion de este archivo es .proto
Ahora tenemos Student.proto como hacemos para pasar eso a Go, pues tenemos la compilación a traves del compilador de buffer que se llama protoC pasara lo que este en Student.proto a un codigo de Go

https://developers.google.com/protocol-buffers/docs/reference/go-generated

Esto es para poder trabajar con los protobufer
$ go get google.golang.org/protobuf

Quick start | Go | gRPC
https://grpc.io/docs/languages/go/quickstart/

Database 
go get github.com/lib/pq

Necesario apra los servicios de grpc, para poder implementar clientes 
y servidores
go get google.golang.org/grpc


Docker database
docker build . -t grpc-db
docker run -p 54321:5432 grpc-db