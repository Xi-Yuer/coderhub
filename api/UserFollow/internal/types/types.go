// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type CreateUserFollowReq struct {
	UserId     int64 `json:"user_id"`     // 关注者ID
	FollowedId int64 `json:"followed_id"` // 被关注者ID
}

type CreateUserFollowResp struct {
	Response
	Data bool `json:"data"` // 操作是否成功
}

type DeleteUserFollowReq struct {
	UserId     int64 `json:"user_id"`     // 取消关注者ID
	FollowedId int64 `json:"followed_id"` // 被取消关注者ID
}

type DeleteUserFollowResp struct {
	Response
	Data bool `json:"data"` // 操作是否成功
}

type GetMutualFollowsReq struct {
	UserId int64 `json:"user_id"` // 用户ID
}

type GetMutualFollowsResp struct {
	Response
	Data UserFollowList `json:"data"`
}

type GetUserFansReq struct {
	UserId   int64 `json:"user_id"`   // 用户ID
	Page     int32 `json:"page"`      // 页码
	PageSize int32 `json:"page_size"` // 每页数量
}

type GetUserFansResp struct {
	Response
	Data UserFollowedList `json:"data"`
}

type GetUserFollowsReq struct {
	UserId   int64 `json:"user_id"`   // 用户ID
	Page     int32 `json:"page"`      // 页码
	PageSize int32 `json:"page_size"` // 每页数量
}

type GetUserFollowsResp struct {
	Response
	Data UserFollowList `json:"data"`
}

type HealthResponse struct {
	Response
}

type IsUserFollowedReq struct {
	UserId     int64 `json:"user_id"`     // 关注者ID
	FollowedId int64 `json:"followed_id"` // 被关注者ID
}

type IsUserFollowedResp struct {
	Response
	Data bool `json:"data"` // 是否已关注
}

type Response struct {
	Code    int32  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
}

type UserFollowInfo struct {
	UserId   int64  `json:"user_id"`  // 用户ID
	Username string `json:"username"` // 用户名
	Avatar   string `json:"avatar"`   // 头像
}

type UserFollowList struct {
	List  []UserFollowInfo `json:"list"`  // 列表
	Total int64            `json:"total"` // 总数
}

type UserFollowedList struct {
	List  []UserFollowInfo `json:"list"`  // 列表
	Total int64            `json:"total"` // 总数
}
