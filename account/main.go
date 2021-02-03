package main

import (
	"log"
	"net"
	"shakebook/account/impl"
	accountpb "shakebook/account/proto/api/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("net listen failed.%v\n", err)
	}
	grpcServer := grpc.NewServer()
	accountpb.RegisterAccountServiceServer(grpcServer, &impl.Server{})
	log.Fatal(grpcServer.Serve(lis))
}
