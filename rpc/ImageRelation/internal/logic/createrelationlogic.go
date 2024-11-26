package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRelationLogic {
	return &CreateRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建图片关系
func (l *CreateRelationLogic) CreateRelation(in *imageRelation.CreateRelationRequest) (*imageRelation.CreateRelationResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.CreateRelationResponse{}, nil
}
