package main

import (
	"context"
	pb "grpc_api/gen/proto"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	if req.Msg == "naber" {
		req.Msg = "iyi senden?"
	}
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {

	go func() {
		mux := runtime.NewServeMux()

		//register
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})

		//server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))
	}()

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Println(err)
	}
}
