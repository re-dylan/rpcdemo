package classroom

import (
	"context"
	logic "rpc/internal/logic/classroom"
	"rpc/internal/svc"
	"rpc/rpc"
)

type Server struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedClassroomServer
}

func NewServer(svcCtx *svc.ServiceContext) *Server {
	return &Server{
		svcCtx: svcCtx,
	}
}

func (s *Server) GetClassroom(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewGetClassroomLogic(ctx, s.svcCtx)
	return l.GetClassroom(in)
}
