package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/controller"
	"bx.com/user-service/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"
	_  "github.com/codyguo/godaemon"
	"bx.com/user-service/utils"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	cfg := config.Parse("config/app.yaml")
	log.Info("[ok] load config")
	bxgo.CreateOrmEngin(cfg.Datasource)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Addr, strconv.Itoa(cfg.Port)))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	//添加拦截器
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		err = utils.Validate.ValidateData(req)
		if err != nil {
			return nil, err
		}
		// 继续处理请求
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	//实例化grpc Server
	//s := grpc.NewServer(opts...)
	s := utils.NewServer()

	//注册API服务
	proto.RegisterUserServiceServer(s, &(controller.UserController{}))
	proto.RegisterTwoFactorServiceServer(s, &(controller.TwoFactorController{}))
	proto.RegisterKycServiceServer(s, &(controller.KycController{}))
	proto.RegisterUserWalletServiceServer(s, &(controller.UserWalletControl{}))
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Infof("[ok] app run at %s:%s", cfg.Addr, strconv.Itoa(cfg.Port))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
