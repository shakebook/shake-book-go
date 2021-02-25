package main

import (
	"context"
	"log"
	"net/http"

	accountpb "shakebook/service/account/proto/api/v1"
	managerpb "shakebook/service/manager/proto/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	startGRPCGateway()
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			EnumsAsInts: true,
			OrigName:    true,
		},
	))
	err := accountpb.RegisterAccountServiceHandlerFromEndpoint(
		c,
		mux,
		":8082",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	err = managerpb.RegisterManagerServiceHandlerFromEndpoint(
		c,
		mux,
		":8084",
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatalf("start grpc gateway failed:%v\n", err)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
