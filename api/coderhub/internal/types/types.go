// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type AcademicList struct {
	Total int64               `json:"total"` // 总数
	List  []AcademicNavigator `json:"list"`  // 学术导航列表
}

type AcademicNavigator struct {
	Id        int64  `json:"id"`         // 学术导航 ID
	UserId    int64  `json:"user_id"`    // 用户 ID
	Education string `json:"education"`  // 学历
	Content   string `json:"content"`    // 内容
	Major     string `json:"major"`      // 专业
	School    string `json:"school"`     // 学校
	WorkExp   string `json:"work_exp"`   // 工作经验
	LikeCount int64  `json:"like_count"` // 点赞数
}

type AddAcademicNavigatorReq struct {
	Education string `json:"education"` // 学历
	Content   string `json:"content"`   // 内容
	Major     string `json:"major"`     // 专业
	School    string `json:"school"`    // 学校
	WorkExp   string `json:"work_exp"`  // 工作经验
}

type AddAcademicNavigatorResp struct {
	Response
	Data bool `json:"data"` // 是否添加成功
}

type CancelLikeAcademicNavigatorReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type CancelLikeAcademicNavigatorResp struct {
	Response
	Data bool `json:"data"` // 是否取消点赞成功
}

type DeleteAcademicNavigatorReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type DeleteAcademicNavigatorResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type DeleteRequest struct {
	ImageId int64 `json:"image_id"` // 图片ID
	UserId  int64 `json:"user_id"`  // 用户ID（用于权限验证）
}

type DeleteResponse struct {
	Response
	Data bool `json:"success"` // 删除是否成功
}

type DeleteUserReq struct {
	Id int64 `path:"id"` // 用户ID
}

type DeleteUserResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type FollowList struct {
	Total int64        `json:"total"` // 总数
	List  []UserFollow `json:"list"`  // 关注列表
}

type FollowUserReq struct {
	FollowUserId int64 `json:"follow_id"` // 被关注用户ID
}

type FollowUserResp struct {
	Response
	Data bool `json:"data"` // 是否关注成功
}

type GetAcademicNavigatorReq struct {
	UserId    int64  `json:"user_id,optional"`   // 用户 ID
	Education string `json:"education,optional"` // 学历
	Content   string `json:"content,optional"`   // 内容
	Major     string `json:"major,optional"`     // 专业
	School    string `json:"school,optional"`    // 学校
	WorkExp   string `json:"work_exp,optional"`  // 工作经验
	Page      int64  `json:"page"`               // 页码
	PageSize  int64  `json:"page_size"`          // 每页大小
}

type GetAcademicNavigatorResp struct {
	Response
	Data *AcademicList `json:"data"` // 学术导航列表
}

type GetFansListReq struct {
	UserId   int64 `form:"user_id"`   // 用户ID
	Page     int64 `form:"page"`      // 页码
	PageSize int64 `form:"page_size"` // 每页数量
}

type GetFansListResp struct {
	Response
	Data FollowList `json:"data"` // 粉丝列表
}

type GetFollowListReq struct {
	UserId   int64 `form:"user_id"`   // 用户ID
	Page     int64 `form:"page"`      // 页码
	PageSize int64 `form:"page_size"` // 每页数量
}

type GetFollowListResp struct {
	Response
	Data FollowList `json:"data"` // 关注列表
}

type GetRequest struct {
	ImageId int64 `path:"image_id"` // 图片ID
}

type GetResponse struct {
	Response
	Data *ImageInfo `json:"data"` // 图片详情
}

type GetUserInfoReq struct {
	Id int64 `path:"id"` // 用户ID
}

type GetUserInfoResp struct {
	Response
	Data *UserInfo `json:"data"` // 用户信息
}

type GetUserListReq struct {
	Page     int64 `form:"page"`      // 页码
	PageSize int64 `form:"page_size"` // 每页数量
}

type GetUserListResp struct {
	Response
	Data UserList `json:"data"` // 用户列表
}

type HealthResp struct {
	Response
	Data bool `json:"data"` // 是否健康
}

type ImageInfo struct {
	ImageId      int64  `json:"image_id"`      // 图片ID
	BucketName   string `json:"bucket_name"`   // MinIO bucket名称
	ObjectName   string `json:"object_name"`   // MinIO中的对象名称
	Url          string `json:"url"`           // 完整的访问URL
	ThumbnailUrl string `json:"thumbnail_url"` // 缩略图URL
	ContentType  string `json:"content_type"`  // 文件MIME类型
	Size         int64  `json:"size"`          // 文件大小(bytes)
	Width        int32  `json:"width"`         // 图片宽度(px)
	Height       int32  `json:"height"`        // 图片高度(px)
	UploadIp     string `json:"upload_ip"`     // 上传者IP
	UserId       int64  `json:"user_id"`       // 上传者ID
	Status       string `json:"status"`        // 状态
	CreatedAt    string `json:"created_at"`    // 创建时间
}

type ImageInfoList struct {
	List  []ImageInfo `json:"images"` // 图片列表
	Total int64       `json:"total"`  // 总数量
}

type ListByUserRequest struct {
	UserId   int64 `json:"user_id"`   // 用户ID
	Page     int32 `json:"page"`      // 页码
	PageSize int32 `json:"page_size"` // 每页数量
}

type ListByUserResponse struct {
	Response
	Data *ImageInfoList `json:"data"` // 图片列表
}

type LoginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type LoginResp struct {
	Response
	Data string `json:"data"` // JWT token
}

type PageRequest struct {
	Page     int32 `form:"page"`      // 页码
	PageSize int32 `form:"page_size"` // 每页数量
}

type PageResponse struct {
	Total    int64 `json:"total"`     // 总数
	Page     int32 `json:"page"`      // 当前页码
	PageSize int32 `json:"page_size"` // 每页数量
}

type PostAcademicNavigatorLikeReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type PostAcademicNavigatorLikeResp struct {
	Response
	Data bool `json:"data"` // 是否点赞成功
}

type RegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Avatar   string `json:"avatar"`   // 头像
	Gender   int32  `json:"gender"`   // 性别 0:未知 1:男 2:女
	Age      int32  `json:"age"`      // 年龄
}

type RegisterResp struct {
	Response
	Data bool `json:"data"` // 是否注册成功
}

type Response struct {
	Code    int32  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
}

type UnfollowUserReq struct {
	FollowUserId int64 `json:"follow_id"` // 被关注用户ID
}

type UnfollowUserResp struct {
	Response
	Data bool `json:"data"` // 是否取消成功
}

type UpdateUserInfoReq struct {
	Id       int64  `path:"id"`       // 用户ID
	Nickname string `json:"nickname"` // 昵称
	Email    string `json:"email"`    // 邮箱
	Phone    string `json:"phone"`    // 手机号
	Avatar   string `json:"avatar"`   // 头像
	Gender   int32  `json:"gender"`   // 性别 0:未知 1:男 2:女
	Age      int32  `json:"age"`      // 年龄
}

type UpdateUserInfoResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}

type UploadResponse struct {
	Response
	Data *ImageInfo `json:"data"` // 上传的图片信息
}

type UserFollow struct {
	Id           int64  `json:"id"`          // 关注ID
	UserId       int64  `json:"user_id"`     // 用户ID
	FollowUserId int64  `json:"follow_id"`   // 被关注用户ID
	Status       int32  `json:"status"`      // 状态 0:未关注 1:已关注
	CreateTime   string `json:"create_time"` // 创建时间
	UpdateTime   string `json:"update_time"` // 更新时间
}

type UserInfo struct {
	Id         int64  `json:"id"`          // 用户ID
	Username   string `json:"username"`    // 用户名
	Password   string `json:"password"`    // 密码
	Nickname   string `json:"nickname"`    // 昵称
	Email      string `json:"email"`       // 邮箱
	Phone      string `json:"phone"`       // 手机号
	Avatar     string `json:"avatar"`      // 头像
	Gender     int32  `json:"gender"`      // 性别 0:未知 1:男 2:女
	Age        int32  `json:"age"`         // 年龄
	Status     int32  `json:"status"`      // 状态 0:正常 1:禁用
	Role       int32  `json:"role"`        // 角色 0:普通用户 1:管理员
	LastLogin  string `json:"last_login"`  // 最后登录时间
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type UserList struct {
	Total int64      `json:"total"` // 总数
	List  []UserInfo `json:"list"`  // 用户列表
}