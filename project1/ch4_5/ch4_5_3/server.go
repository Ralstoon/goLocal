package main

import (
	"context"
	"fmt"
	pc "go_code/project1/protocCode"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type Authentication struct {
	Login    string
	Password string
}

func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"login": a.Login, "password": a.Password}, nil
}

func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

var (
	port = ":5000"
)

type myGrpcServer struct{}

func (s *myGrpcServer) SayHello(ctx context.Context, in *pc.HelloRequest) (*pc.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing authenication")
	}

	var (
		appid  string
		appkey string
	)
	fmt.Println(md)
	if val, ok := md["login"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != "wuyuhang" || appkey != "123456" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token: appid=%s, appkey=%s", appid, appkey)
	}

	return &pc.HelloReply{Message: "Hello:" + in.GetName()}, nil
}

func startServer() {
	server := grpc.NewServer()
	pc.RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not listen on :%s:%s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc server error: %s", err)
	}
}

func goClientWork() {
	auth := Authentication{
		Login:    "wuyuhang",
		Password: "123456",
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pc.NewGreeterClient(conn)

	reply, err := client.SayHello(context.Background(), &pc.HelloRequest{Name: "wuyuhang1233"})
	if err != nil {
		log.Fatal("Could not Greet:%v", err)
	}

	log.Printf("Finished ClientWork: %s", reply.Message)
}

func main() {
	go startServer()
	time.Sleep(time.Second * 2)

	goClientWork()
}
