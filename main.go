package main

import (
	"fmt"
	"net"
	"os"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"github.com/tap2joy/CenterService/server"
	pb "github.com/tap2joy/Protocols/go/grpc/center"
)

func main() {
	lis, err := net.Listen("tcp", ":9100")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		os.Exit(1)
	}
	fmt.Println("CenterService start at port: 9100")

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		grpc_validator.UnaryServerInterceptor())))

	pb.RegisterCenterServiceServer(s, &server.Server{})
	grpc_health_v1.RegisterHealthServer(s, &server.HealthServer{})
	reflection.Register(s)
	s.Serve(lis)
}
