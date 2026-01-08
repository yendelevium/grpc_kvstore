package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/yendelevium/grpc_kvstore/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedKVStoreServer
	mp map[string]string
}

func (s *server) Put(_ context.Context, in *pb.PutArgs) (*pb.PutResponse, error) {
	// Add key-value to the store
	log.Println("Recieved PUT request")
	s.mp[in.Key] = in.Value
	return &pb.PutResponse{
		Value: s.mp[in.Key],
	}, nil
}

func (s *server) Get(_ context.Context, in *pb.GetArgs) (*pb.GetResponse, error) {
	log.Println("Recieved GET request")
	// Check for key in store
	val, ok := s.mp[in.Key]
	if !ok {
		// If value not in store, return error
		return &pb.GetResponse{}, fmt.Errorf("key '%s' doens't exist in the store", in.Key)
	}

	// Return value
	return &pb.GetResponse{
		Value: val,
	}, nil
}

func main() {
	// General creating the server stuff
	lis, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Flopped the listen, %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKVStoreServer(s, &server{
		mp: make(map[string]string),
	})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
