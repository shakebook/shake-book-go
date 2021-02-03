package main

import (
	"context"
	"log"
	"net/http"

	accountpb "shakebook/account/proto/api/v1"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	startGRPCGateway()
	// conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("grpc dial failed: %v\n", err)
	// }
	// defer conn.Close()

	// client := accountpb.NewAccountServiceClient(conn)
	// res, err := client.GetAccount(context.Background(), &accountpb.GetAccountRequest{
	// 	Id: 1,
	// })
	// if err != nil {
	// 	log.Printf("get account failed:%v\n", err)
	// 	return
	// }
	// log.Printf("account:%s\n", res)
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
		":8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalf("start grpc gateway failed:%v\n", err)
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
