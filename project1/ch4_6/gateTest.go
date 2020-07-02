package main

import (
	"context"
	pc "go_code/project1/protocCode"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	port = ":5000"
)

type MyRestServer struct{}

// type RestServiceServer interface {
// 	Get(context.Context, *StringMessage) (*StringMessage, error)
// 	Post(context.Context, *StringMessage) (*StringMessage, error)
// }

func (s *MyRestServer) Get(ctx context.Context, in *pc.StringMessage) (*pc.StringMessage, error) {
	return &pc.StringMessage{Value: "Get in " + in.Value}, nil
}
func (s *MyRestServer) Post(ctx context.Context, in *pc.StringMessage) (*pc.StringMessage, error) {
	return &pc.StringMessage{Value: "Post in " + in.Value}, nil
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	input := []grpc.DialOption{grpc.WithInsecure()}
	err := pc.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", input)
	if err != nil {
		log.Fatal(err)
	}
	return http.ListenAndServe(":8080", mux)
}

func startRestServer() {
	server := grpc.NewServer()
	pc.RegisterRestServiceServer(server, new(MyRestServer))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}

func main() {

	go startRestServer()

	if err := run(); err != nil {
		log.Panicln(err)
	}
}
