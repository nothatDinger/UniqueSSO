package router

import (
	"github.com/UniqueStudio/UniqueSSO/controller"
	"github.com/UniqueStudio/UniqueSSO/pb/sso"
	"google.golang.org/grpc"
)

func InitRPC(s *grpc.Server) {
	sso.RegisterSSOServiceServer(s, controller.NewSSOServiceServer())
}
