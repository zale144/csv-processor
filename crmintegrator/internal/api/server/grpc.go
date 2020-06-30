package server

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"crmintegrator/internal/app/user"
	"crmintegrator/internal/config"

	"crmintegrator/internal/api/handler"
	"crmintegrator/internal/pb"
)

func Start(cfg *config.Config, userSrvc user.UserService) error {

	lis, err := net.Listen("tcp", cfg.GRPCURL)
	if err != nil {
		return errors.Wrap(err, "failed to listen gRPC")
	}

	grpcServer := grpc.NewServer()

	userHandler := handler.NewUser(userSrvc)
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Printf("gRPC listening on: %s", cfg.GRPCURL)
	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve gRPC")
	}
	return nil
}
