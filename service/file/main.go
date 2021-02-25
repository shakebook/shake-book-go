package main

import (
	"log"
	"net"
	"shakebook/common/auth"
	"shakebook/common/conn"
	lg "shakebook/common/log"
	"shakebook/service/account/dao"
	"shakebook/service/account/impl"
	accountpb "shakebook/service/account/proto/api/v1"

	authpb "shakebook/service/auth/proto/api/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := lg.NewZapLogger()
	lis, err := net.Listen("tcp", ":8083")
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

	accountpb.RegisterAccountServiceServer(grpcServer, &impl.Server{
		Logger:     logger,
		AuthClient: authpb.NewAuthServiceClient(authConn),
		Dao: &dao.Conn{
			Mysql: &conn.Mysql{
				UserName:     "root",
				Password:     "12345678",
				Addr:         "localhost:3306",
				DatabaseName: "shakebook",
			},
		},
	})
	logger.Sugar().Fatal(grpcServer.Serve(lis))
}
