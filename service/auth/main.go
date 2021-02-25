package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	lg "shakebook/common/log"
	"shakebook/service/auth/impl"
	authpb "shakebook/service/auth/proto/api/v1"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := lg.NewZapLogger()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("net listen failed.%v\n", err)
	}
	file, err := os.Open("./private_key")
	if err != nil {
		logger.Fatal("open private_key failed.", zap.Error(err))
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Fatal("ioutil read all private_key failed.", zap.Error(err))

	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(fileBytes)
	if err != nil {
		logger.Fatal("cannot parse private key", zap.Error(err))
	}
	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, impl.NewServer(
		privKey,
		"yangjiafeng",
		time.Minute*300,
	))
	log.Fatal(grpcServer.Serve(lis))
}
