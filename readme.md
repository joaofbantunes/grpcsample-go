# About
Stupid simple sample to try gRPC and Go (blatantly ripped from the official [helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld) example).

This will be used as a server (although there's a client here for testing), and there's a C# client [here](https://github.com/joaofbantunes/grpcsample-dotnet).

Consider I'm on Windows for the next steps

# Steps I took (setup)
1. Installed Protobuf compiler (protoc)
2. Installed Go :)
3. Set GOPATH environment variable to where the projects will be kept
4. Installed gRPC for Go `go get google.golang.org/grpc`
5. Installed the Go code generator plugin `go get -u github.com/golang/protobuf/protoc-gen-go` (it will be on the GOPATH and protoc is able to find it)


# Steps I took (development)
1. Created the project folder under src, under GOPATH
2. Created the service definition (servicedef folder), consisting on the messages to pass (IncrementRequest and IncrementResponse) and the available operations (service Counter.Increment)
3. Generated the Go files from the proto definition running the following command on the root of the project  `protoc --go_out=plugins=grpc:generated servicedef/*.proto`
4. Installed the generated package on my GOPATH by running the following command (still on the project root) `go install ./generated/servicedef`
5. Implemented the server on server\main.go and the client on client\main.go
6. On a command window ran the server `go run server\main.go`
7. On another command window ran the client `go run client\main.go`
8. Stuff works, YAY!

# Notes
Followed parts of the [gRPC Go Quick Start](http://www.grpc.io/docs/quickstart/go.html) and the [helloworld](https://github.com/grpc/grpc-go/tree/master/examples/helloworld) example.