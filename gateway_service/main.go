package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	pb "hellogrpc/api"

	"google.golang.org/grpc"
)

const (
	contractAddress = "localhost:50051"
	timeoutSec      = 10
)

type createContractRequest struct {
	ContractType string `json:"contractType"`
}

type createContractReply struct {
	ContractID string `json:"contractId"`
}

func handleContract(w http.ResponseWriter, r *http.Request, c pb.ContractClient) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reportErr(err, w)
		return
	}

	req := createContractRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		reportErr(err, w)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSec*time.Second)
	defer cancel()
	reply, err := c.CreateContract(ctx, &pb.CreateContractRequest{ContractType: req.ContractType})
	if err != nil {
		reportErr(err, w)
		return
	}

	rep := createContractReply{ContractID: reply.ContractId}
	repBody, err := json.Marshal(rep)
	if err != nil {
		reportErr(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", string(repBody))
}

func reportErr(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusExpectationFailed)
	log.Printf("%v", err)
	fmt.Fprintf(w, "%v", err)
}

func main() {
	log.Print("Starting Gateway service..")

	conn, err := grpc.Dial(contractAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	contractClient := pb.NewContractClient(conn)

	http.HandleFunc("/contracts", func(w http.ResponseWriter, r *http.Request) {
		handleContract(w, r, contractClient)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("failed to start server, %v", err)
	}
}
