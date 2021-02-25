package main

import (
	"context"
	"log"
	"net"
	"shakebook/common/auth"
	"shakebook/common/conn"
	lg "shakebook/common/log"

	"shakebook/service/manager/dao"
	"shakebook/service/manager/impl"
	namagerpb "shakebook/service/manager/proto/api/v1"

	authpb "shakebook/service/auth/proto/api/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := lg.NewZapLogger()
	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("net listen failed.%v\n", err)
	}

	authConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		logger.Fatal("cannot connect auth service", zap.Error(err))
	}

	var opts []grpc.ServerOption
	interceptor, err := auth.NewInterceptor("./public_key")
	if err != nil {
		logger.Fatal("create interceptor failed", zap.Error(err))
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	grpcServer := grpc.NewServer(opts...)

	namagerpb.RegisterManagerServiceServer(grpcServer, &impl.Server{
		Logger:     logger,
		AuthClient: authpb.NewAuthServiceClient(authConn),
		Dao: &dao.Conn{
			Mysql: &conn.Mysql{
				UserName:     "root",
				Password:     "12345678",
				Addr:         "localhost:3306",
				DatabaseName: "shakebook",
			},
			Redis: &conn.Redis{
				Addr:     "localhost:6379",
				Password: "12345678",
				Context:  context.Background(),
			},
		},
	})
	logger.Sugar().Fatal(grpcServer.Serve(lis))
}
