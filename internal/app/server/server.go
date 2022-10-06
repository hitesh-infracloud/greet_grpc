package server

import (
	"context"
	"fmt"
	greetpb "github.com/hitesh-infracloud/awesome-code/greet_app/internal/pkg/proto"
	"log"
)

type server struct {
}

func GetNewGreetServer() *server {
	return &server{}
}

func (*server) Greet(_ context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	log.Printf("Greet function was invoked with req: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := fmt.Sprintf("Hello, %s %s", firstName, lastName)
	return &greetpb.GreetingResponse{
		Result: result,
	}, nil
}
