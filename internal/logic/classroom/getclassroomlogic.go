package classroom

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc/internal/svc"
	"rpc/rpc"
)

type GetClassroomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassroomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassroomLogic {
	return &GetClassroomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetClassroomLogic) GetClassroom(in *rpc.Request) (*rpc.Response, error) {
	// todo: add your logic here and delete this line

	return &rpc.Response{}, nil
}
