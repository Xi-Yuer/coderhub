package logic

import (
	"context"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetUserByIDLogic {
	return &BatchGetUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchGetUserByID 批量获取用户信息
func (l *BatchGetUserByIDLogic) BatchGetUserByID(in *user.BatchGetUserByIDRequest) (*user.BatchGetUserByIDResponse, error) {
	users, err := l.svcCtx.UserRepository.BatchGetUserByID(in.UserIds)
	if err != nil {
		return nil, err
	}

	userInfos := make([]*user.UserInfo, len(users))
	for i, val := range users {
		userInfos[i] = &user.UserInfo{
			UserId:    val.ID,
			UserName:  val.UserName,
			Avatar:    val.Avatar.String,
			Email:     val.Email.String,
			NickName:  val.NickName.String,
			IsAdmin:   val.IsAdmin,
			Status:    val.Status,
			CreatedAt: val.CreatedAt.Unix(),
			UpdatedAt: val.UpdatedAt.Unix(),
		}
	}

	return &user.BatchGetUserByIDResponse{UserInfos: userInfos}, nil
}
