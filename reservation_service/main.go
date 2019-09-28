package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "hellogrpc/api"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct{}

func (s *server) Reserve(ctx context.Context, in *pb.ReservationRequest) (*pb.ReservationReply, error) {
	if ctx.Err() != nil {
		log.Printf("%v", ctx.Err())
		return nil, ctx.Err()
	}

	log.Printf("Reserve: %v", in.GetQuantity())
	reservationID := uuid.New().String()

	time.Sleep(time.Duration(in.GetQuantity()) * time.Second)

	if ctx.Err() != nil {
		log.Printf("%v", ctx.Err())
		return nil, ctx.Err()
	}

	return &pb.ReservationReply{ReservationId: reservationID}, nil
}

func main() {
	log.Print("Starting Reservation service..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReservationServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
