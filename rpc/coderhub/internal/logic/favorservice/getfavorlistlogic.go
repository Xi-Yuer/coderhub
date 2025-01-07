package favorservicelogic

import (
	"coderhub/model"
	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/svc"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavorListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavorListLogic {
	return &GetFavorListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFavorList 获取收藏夹列表
func (l *GetFavorListLogic) GetFavorList(in *coderhub.GetFavorListRequest) (*coderhub.GetFavorListResponse, error) {
	list, total, err := l.svcCtx.UserFavorEntityRepository.GetList(l.ctx, &model.UserFavor{
		UserId:      in.UserId,
		FavorFoldId: in.FavorFolderId,
		EntityType:  in.EntityType,
	}, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}
	favors := make([]*coderhub.Favor, 0, len(list))
	for _, v := range list {
		favor := &coderhub.Favor{
			Id:            int64(v.ID),
			UserId:        v.UserId,
			FavorFolderId: v.FavorFoldId,
			EntityId:      v.EntityId,
			EntityType:    v.EntityType,
			CreateTime:    v.CreatedAt.Unix(),
		}
		favors = append(favors, favor)
	}
	fmt.Println("len(favors)", len(favors))

	return &coderhub.GetFavorListResponse{
		Favors: favors,
		Total:  total,
	}, nil
}
