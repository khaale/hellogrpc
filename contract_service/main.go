package main

import (
	"context"
	"log"
	"net"

	pb "hellogrpc/api"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const (
	port               = ":50051"
	reservationAddress = "localhost:50052"
)

type server struct {
	ReservationClient pb.ReservationClient
}

func (s *server) CreateContract(ctx context.Context, in *pb.CreateContractRequest) (*pb.CreateContractReply, error) {
	if ctx.Err() != nil {
		log.Printf("%v", ctx.Err())
		return nil, ctx.Err()
	}

	log.Printf("Received: %v", in.ContractType)
	contractID := uuid.New().String()

	if in.ContractType == "complex" {
		_, err := s.ReservationClient.Reserve(ctx, &pb.ReservationRequest{Quantity: 6})
		if err != nil {
			log.Printf("%v", ctx.Err())
			return nil, ctx.Err()
		}
		_, err = s.ReservationClient.Reserve(ctx, &pb.ReservationRequest{Quantity: 60})
		if err != nil {
			log.Printf("%v", ctx.Err())
			return nil, ctx.Err()
		}
	} else {
		_, err := s.ReservationClient.Reserve(ctx, &pb.ReservationRequest{Quantity: 1})
		if err != nil {
			log.Printf("%v", ctx.Err())
			return nil, ctx.Err()
		}
	}

	return &pb.CreateContractReply{ContractId: contractID}, nil
}

func main() {
	log.Print("Starting Contract service..")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := grpc.Dial(reservationAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	s := grpc.NewServer()
	pb.RegisterContractServer(s, &server{ReservationClient: pb.NewReservationClient(conn)})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
