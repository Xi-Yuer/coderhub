package favorfoldservicelogic

import (
	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/svc"
	"coderhub/shared/utils"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavorFoldListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavorFoldListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavorFoldListLogic {
	return &GetFavorFoldListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFavorFoldList 获取列表
func (l *GetFavorFoldListLogic) GetFavorFoldList(in *coderhub.GetFavorFoldListRequest) (*coderhub.GetFavorFoldListResponse, error) {
	logx.Infof("获取收藏夹列表: %+v", in)
	userId, err := utils.GetUserMetaData(l.ctx)
	if err != nil {
		return nil, err
	}

	formatInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	list, total, err := l.svcCtx.UserFavorFolderRepository.GetList(l.ctx, in.UserId, formatInt, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	li := make([]*coderhub.FavorFold, 0, len(list))
	for _, v := range list {
		li = append(li, &coderhub.FavorFold{
			Id:          int64(v.ID),
			Name:        v.FavorName,
			Description: v.Description,
			IsPublic:    v.IsPublic,
			UserId:      v.UserId,
			CreateTime:  v.CreatedAt.Unix(),
			UpdateTime:  v.UpdatedAt.Unix(),
		})
	}

	return &coderhub.GetFavorFoldListResponse{
		FavorFolds: li,
		Total:      total,
	}, nil
}
