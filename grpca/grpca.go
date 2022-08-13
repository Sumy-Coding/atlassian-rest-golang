package grpca

import (
	pb "confluence-rest-golang/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

//https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type MyServer struct {
	pb.PageServiceServer
}

func (s *MyServer) GetPage(ctx context.Context, in *pb.PageRequest) (*pb.Content, error) {
	//return &com_andmal.PageResponse{Message: "Hello again " + in.GetName()}, nil
	return &pb.Content{
		Id: "77777777777",
	}, nil
}

func (srv MyServer) GetPages(ctx context.Context, req *pb.PageRequest) (*[]pb.Content, error) {
	//
}

func (srv *MyServer) InitServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPageServiceServer(s, &MyServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
