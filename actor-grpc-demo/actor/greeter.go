package actor

import (
	"fmt"

	pb "actor-grpc-demo/proto"

	"github.com/asynkron/protoactor-go/actor"
)

type GreeterActor struct{}

func (g *GreeterActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *pb.GreetRequest:
		reply := &pb.GreetResponse{
			Message: fmt.Sprintf("Hello, %s!", msg.Name),
		}
		ctx.Respond(reply)
	}
}

func NewGreeterActor() actor.Actor {
	return &GreeterActor{}
}
