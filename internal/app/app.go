package app

import (
	"fmt"
	"github.com/hitesh-infracloud/awesome-code/greet_app/internal/app/server"
	greetpb "github.com/hitesh-infracloud/awesome-code/greet_app/internal/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type App struct {
	grpcServer *grpc.Server
}

func GetNewGreetApp() *App {
	return &App{}
}

func (a *App) Start() {
	fmt.Println("Starting greet server on port 50059....")

	servAddr := ":50059"
	if len(os.Getenv("GRPC_SRV_ADDR")) > 0 {
		servAddr = os.Getenv("GRPC_SRV_ADDR")
	}
	lis, err := net.Listen("tcp", servAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	a.grpcServer = grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(a.grpcServer, server.GetNewGreetServer())

	go func() {
		if err := a.grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
}

func (a *App) Shutdown() {
	a.grpcServer.GracefulStop()
}
