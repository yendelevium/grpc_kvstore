package main

import (
	"context"
	"log"
	"time"

	pb "github.com/yendelevium/grpc_kvstore/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewKVStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// put key; response isn't needed
	_, err = c.Put(ctx, &pb.PutArgs{
		Key:   "kv",
		Value: "store",
	})
	if err != nil {
		log.Printf("Error in gRPC call: %v", err)
	}

	// Existing Key
	resp, err := c.Get(ctx, &pb.GetArgs{
		Key: "kv",
	})
	if err != nil {
		log.Printf("Error in gRPC call: %v", err)
	}
	log.Printf("The value for key '%s' is '%s'", "kv", resp.Value)

	// Non-existing key
	resp, err = c.Get(ctx, &pb.GetArgs{
		Key: "randomkey",
	})
	if err != nil {
		log.Printf("Error in gRPC call: %v", err)
	}

}
