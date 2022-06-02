package user

import (
	"context"

	"rpc/internal/logic/user"
	"rpc/internal/svc"
	"rpc/rpc"
)

type Server struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedUserServer
}

func NewServer(svcCtx *svc.ServiceContext) *Server {
	return &Server{
		svcCtx: svcCtx,
	}
}

func (s *Server) GetUserInfo(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := user.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
