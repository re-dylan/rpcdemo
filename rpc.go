package main

import (
	"flag"
	"fmt"
	"rpc/internal/server/classroom"
	"rpc/internal/server/user"

	"rpc/internal/config"
	"rpc/internal/server"
	"rpc/internal/svc"
	"rpc/rpc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/rpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		rpc.RegisterRpcServer(grpcServer, server.NewRpcServer(ctx))
		rpc.RegisterUserServer(grpcServer, user.NewServer(ctx))
		rpc.RegisterClassroomServer(grpcServer, classroom.NewServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	rpc.NewUserClient()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
