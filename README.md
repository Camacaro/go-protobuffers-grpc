# Protobuffers y gRPC
Nos ayuda a crear microservicios de manera más rápida y de alto rendimiento.

## Protobuffers
Es un protocolo de buffers (algo así como json), es un serializado de datos como XML, nos permite armar estructura de datos. 

### Message
Nos permite definir un mensage. Vamos a crear un mensaje o una structura nueva Nombre {} -> dentro de las llaves creamos la estructua de datos, tipo de la variable (string) y el nombre de la variable y se le asigna un valor que tiene que ser unico, ademas tiene que ser secuencial con las demas variables (estructua de dato), sera la posicion donde será serializado

message Student {
  string id = 1
  string name = 2  
}

### Seralizar -> Deserializar 
Ya que como en el protobuffer se encuentran estructurada este proceso de deserialización será más rápida.


la extencion de este archivo es .proto ahora tenemos Student.proto. ¿Cómo hacemos para pasar eso a Go?, pues tenemos la compilación a traves del compilador de buffer que se llama ```protoC``` pasara lo que este en Student.proto a un [codigo de Go](https://developers.google.com/protocol-buffers/docs/reference/go-generated)

## Paquetes necesario
* Esto es para poder trabajar con los protobufer ```go get google.golang.org/protobuf```
* Necesario para los servicios de grpc, para poder implementar clientes y servidores ```go get google.golang.org/grpc```
* Database Postgres ```go get github.com/lib/pq```

## Quick start | Go | gRPC
https://grpc.io/docs/languages/go/quickstart/


## Docker database 
Dentro de la carpeta ```database``` generar la imagen y levantarla
```
docker build . -t grpc-db
docker run -p 54321:5432 grpc-db
```

## Proyecto
### Implementando Strameaming del lado del cliente
Enviaremos una serie de preguntar al servidor para que sean almacenadas y nos devuelva una sola respuesta

### Implementando el Strameaming del lado del servidor 
Se envia un dato, y el servidor empieza a transmitir datos hasta que se acabe de enviarlos todos 

### Implementacion Strameaming bidireccional
Tanto como el cliente como el servidor van a enviar y recibir data al mismo tiempo. 

## Limitaciones de gRPC 
Se complican si queremos conectar gRPC en una aplicacion web ya que  no hay una forma nativa de implementarolo pero existe una  manera, crear un Proxy Rest, el problema es que agremagos una capa extra a nuestro proyecto.

* [grpc-web](https://github.com/grpc/grpc-web)
* [envoy](https://www.envoyproxy.io)

## Tips
* RPC: Remote Procedure Call
* ¿Qué nos permite RPC? Ejecutar código de un servidor como si hubiera sido implementado en los clientes.
* Que es gRPC: Un Framework de alto rendimiento para implementaciones de RPC.
* Protobuffers:  Nos permite definir messages y estructuras de datos fuertemente tipadas.
* ventaja de ProtoBuffer sobre JSON: Mejores velocidades de serialización/deserialización.
* ventaja de JSON sobre ProtoBuffer: Más fácil de leer a nivel humano.
* **Método Unary/Unario de gRPC**: Similar al fomato request/response en el cuál la interacción ocurre de una a una.
* **método Client Streaming/Streaming del Lado del Cliente**: El cliente envía la data a través de un streaming y el servidor va a responder una sola vez.
* **método Server Streaming/Streaming del lado del Servidor**: El cliente envía una sola vez la data y el servidor responde utilizando un streaming.
* **método Bidirectional Streaming/Streaming bidireccional:** Ambos, cliente y servidor se comunican utilizando streaming.
* **alternativas para implementar gRPC directamente en el navegador:** El proyecto grpc-web y un proxy grpc/rest.
* **Dos grandes innovaciones de gRPC:** HTTP2 y ProtoBuffers
* **compilador de ProtoBuffer:** protoc
* **compilación de un archivo .proto con las opciones de Go:**: Un paquete de Go
* **servicio Unary/Unario:** ```rpc Nombre(Request) returns (Response)```
* **servicio con Client Streaming/Streaming del lado del Cliente:** ```rpc Nombre(stream Request) returns (Response)```
* **servicio con Server Streaming/Streaming del lado del Servidor:** ```rpc Nombre(Request) returns (stream Response)```
* **servicio con Bidirectional Streaming/Streaming Bidireccional:** ```rpc Nombre(stream Request) returns (stream Response)```