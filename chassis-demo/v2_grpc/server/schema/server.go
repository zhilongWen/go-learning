package schema

import (
	pb "chassis-demo/v2_grpc/helloworld"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	fmt.Println("server say hello...")

	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md["x-user"])

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
