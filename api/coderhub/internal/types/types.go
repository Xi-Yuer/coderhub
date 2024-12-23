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

type Article struct {
	Id           int64    `json:"id" form:"id"`                     // 主键 ID
	Type         string   `json:"type" form:"type"`                 // 内容类型：长文或短文
	Title        string   `json:"title" form:"title"`               // 标题
	Content      string   `json:"content" form:"content"`           // 内容
	Summary      string   `json:"summary" form:"summary"`           // 摘要
	ImageUrls    []string `json:"imageUrls" form:"imageUrls"`       // 图片 URL 列表
	CoverImage   *string  `json:"coverImage" form:"coverImage"`     // 封面图片 URL
	AuthorId     int64    `json:"authorId" form:"authorId"`         // 作者 ID
	Tags         []string `json:"tags" form:"tags"`                 // 标签列表
	ViewCount    int64    `json:"viewCount" form:"viewCount"`       // 阅读次数
	LikeCount    int64    `json:"likeCount" form:"likeCount"`       // 点赞次数
	CommentCount int64    `json:"commentCount" form:"commentCount"` // 评论数
	Status       string   `json:"status" form:"status"`             // 文章状态
	CreatedAt    int64    `json:"createdAt" form:"createdAt"`       // 创建时间
	UpdatedAt    int64    `json:"updatedAt" form:"updatedAt"`       // 更新时间
}

type CancelLikeAcademicNavigatorReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type CancelLikeAcademicNavigatorResp struct {
	Response
	Data bool `json:"data"` // 是否取消点赞成功
}

type Comment struct {
	Id              int64       `json:"id"`                 // 评论ID
	EntityID        int64       `json:"entity_id"`          // 文章ID
	Content         string      `json:"content"`            // 评论内容
	RootId          int64       `json:"root_id"`            // 根评论ID
	ParentId        int64       `json:"parent_id"`          // 父评论ID
	UserInfo        *UserInfo   `json:"user_info"`          // 评论者信息
	CreatedAt       int64       `json:"created_at"`         // 创建时间
	UpdatedAt       int64       `json:"updated_at"`         // 更新时间
	Replies         []*Comment  `json:"replies"`            // 子评论列表
	ReplyToUserInfo *UserInfo   `json:"reply_to_user_info"` // 被回复者信息
	RepliesCount    int64       `json:"replies_count"`      // 子评论数量
	LikeCount       int32       `json:"like_count"`         // 点赞数
	Images          []ImageInfo `json:"images"`             // 评论图片列表
}

type CreateArticleReq struct {
	Type         string   `json:"type,options=article|micro_post"` // 内容类型
	Title        string   `json:"title"`                           // 标题
	Content      string   `json:"content"`                         // 内容
	Summary      string   `json:"summary"`                         // 摘要
	ImageIds     []int64  `json:"imageIds"`                        // 图片 URL 列表
	CoverImageID int64    `json:"coverImageID"`                    // 封面图片 URL
	Tags         []string `json:"tags"`                            // 标签列表
	Status       string   `json:"status,options=draft|published"`  // 文章状态
}

type CreateArticleResp struct {
	Response
	Data int64 `json:"data"` // 文章详情
}

type CreateCommentReq struct {
	EntityID   int64   `json:"entity_id"`    // 文章ID
	Content    string  `json:"content"`      // 评论内容
	RootId     int64   `json:"root_id"`      // 根评论ID
	ParentId   int64   `json:"parent_id"`    // 父评论ID（可选）
	ReplyToUID int64   `json:"reply_to_uid"` // 回复的目标评论ID（可选）
	ImageIds   []int64 `json:"image_ids"`    // 图片ID列表
}

type CreateCommentResp struct {
	Response
	Data *Comment `json:"data"` // 创建的评论
}

type CreateQuestionBankReq struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty,options=default|easy|medium|hard"`
	Tags        []string `json:"tags"`
	CoverImage  int64    `json:"coverImage"` // 封面图片 URL
}

type CreateQuestionBankResp struct {
	Response
	Data bool `json:"data"` // 题库详情
}

type CreateQuestionReq struct {
	Title      string `json:"title"`
	BankId     int64  `json:"bankId"`
	Content    string `json:"content"`
	Difficulty string `json:"difficulty,options=default|easy|medium|hard"`
}

type CreateQuestionResp struct {
	Response
	Data bool `json:"data"` // 题目详情
}

type DeleteAcademicNavigatorReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type DeleteAcademicNavigatorResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type DeleteArticleReq struct {
	Id int64 `path:"id"` // 文章 ID
}

type DeleteArticleResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type DeleteCommentReq struct {
	CommentId int64 `path:"comment_id"` // 评论ID
}

type DeleteCommentResp struct {
	Response
	Data bool `json:"data"` // 删除是否成功
}

type DeleteQuestionBankReq struct {
	Id int64 `path:"id"` // 题库 ID
}

type DeleteQuestionBankResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type DeleteQuestionReq struct {
	Id int64 `path:"id"` // 题目 ID
}

type DeleteQuestionResp struct {
	Response
	Data bool `json:"data"` // 是否删除成功
}

type DeleteRequest struct {
	ImageId int64 `json:"image_id"` // 图片ID
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
	Total int64      `json:"total"` // 总数
	List  []UserInfo `json:"list"`  // 关注列表
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

type GetArticleReq struct {
	Id int64 `path:"id"` // 文章 ID
}

type GetArticleResp struct {
	Response
	Data *Article `json:"data"` // 文章详情
}

type GetCommentRepliesReq struct {
	CommentId int64 `path:"comment_id"` // 评论ID
	Page      int32 `form:"page"`       // 页码
	PageSize  int32 `form:"page_size"`  // 每页数量
}

type GetCommentRepliesResp struct {
	Response
	Data List `json:"data"` // 子评论列表
}

type GetCommentReq struct {
	CommentId int64 `path:"comment_id"` // 评论ID
}

type GetCommentResp struct {
	Response
	Data *Comment `json:"data"` // 评论详情
}

type GetCommentsReq struct {
	EntityID int64 `path:"entity_id"` // 文章ID
	Page     int32 `form:"page"`      // 页码
	PageSize int32 `form:"page_size"` // 每页数量
}

type GetCommentsResp struct {
	Response
	Data List `json:"data"` // 评论列表
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

type GetQuestionBankListReq struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

type GetQuestionBankListResp struct {
	Response
	Data QuestionBankList `json:"data"` // 题库列表
}

type GetQuestionBankReq struct {
	Id int64 `path:"id"` // 题库 ID
}

type GetQuestionBankResp struct {
	Response
	Data *Question `json:"data"` // 题库详情
}

type GetQuestionListReq struct {
	BankId   int64 `json:"bankId"`
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

type GetQuestionListResp struct {
	Response
	Data QuestionList `json:"data"` // 题目列表
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
	CreatedAt    int64  `json:"created_at"`    // 创建时间
}

type ImageInfoList struct {
	List  []ImageInfo `json:"images"` // 图片列表
	Total int64       `json:"total"`  // 总数量
}

type List struct {
	List  []*Comment `json:"list"`  // 评论列表
	Total int32      `json:"total"` // 总评论数
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

type Question struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	BankId    int64  `json:"bankId"`
	Content   string `json:"content"`
	Difficult string `json:"difficult"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type QuestionBank struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Difficulty  string     `json:"difficulty"`
	Tags        []string   `json:"tags"`
	CoverImage  *ImageInfo `json:"coverImage" form:"coverImage"`
	CreateUser  *UserInfo  `json:"createUser"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
}

type QuestionBankList struct {
	Total int64           `json:"total"`
	List  []*QuestionBank `json:"list"`
}

type QuestionList struct {
	Total int64            `json:"total"`
	List  []*QuestionMenus `json:"list"`
}

type QuestionMenus struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type RegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type RegisterResp struct {
	Response
	Data bool `json:"data"` // 是否注册成功
}

type ResetPasswordByLinkReq struct {
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required,min=8,max=32"`
	ConfirmPassword string `form:"confirmPassword" validate:"required,eqfield=Password"`
	Token           string `form:"token" validate:"required"`
}

type ResetPasswordByLinkResp struct {
	Response
	Data bool `json:"data"` // 是否重置成功
}

type Response struct {
	Code    int32  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
}

type SendResetPasswordLinkReq struct {
	Email string `form:"email"` // 邮箱
}

type SendResetPasswordLinkResp struct {
	Response
	Data bool `json:"data"` // 是否发送成功
}

type UnfollowUserReq struct {
	FollowUserId int64 `json:"follow_id"` // 被关注用户ID
}

type UnfollowUserResp struct {
	Response
	Data bool `json:"data"` // 是否取消成功
}

type UpdateArticleReq struct {
	Id           int64    `path:"id"`                        // 文章 ID
	Title        string   `json:"title,optional"`            // 标题
	Content      string   `json:"content,optional,optional"` // 内容
	Summary      string   `json:"summary,optional"`          // 摘要
	ImageIds     []int64  `json:"imageIds,optional"`         // 图片 URL 列表
	CoverImageID int64    `json:"coverImageID,optional"`     // 封面图片 URL
	Tags         []string `json:"tags,optional"`             // 标签列表
	Status       string   `json:"status,optional"`           // 文章状态
}

type UpdateArticleResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}

type UpdateCommentLikeCountReq struct {
	CommentId int64 `json:"comment_id"` // 评论ID
}

type UpdateCommentLikeCountResp struct {
	Response
	Data bool `json:"data"` // 更新是否成功
}

type UpdateLikeCountReq struct {
	Id int64 `json:"id"` // 文章 ID
}

type UpdateLikeCountResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}

type UpdatePasswordReq struct {
	OldPassword string `json:"old_password"` // 旧密码
	NewPassword string `json:"new_password"` // 新密码
}

type UpdatePasswordResp struct {
	Response
	Data bool `json:"data"` // 是否修改成功
}

type UpdateUserAvatarReq struct {
	Avatar int64 `json:"avatar"` // 头像
}

type UpdateUserAvatarResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}

type UpdateUserInfoReq struct {
	Id       int64  `path:"id"`                          // 用户ID
	Nickname string `json:"nickname,optional"`           // 昵称
	Email    string `json:"email,optional"`              // 邮箱
	Phone    string `json:"phone,optional"`              // 手机号
	Gender   int32  `json:"gender,options=0|1,optional"` // 性别 0:未知 1:男 2:女
	Age      int32  `json:"age,range=[0:120],optional"`  // 年龄
}

type UpdateUserInfoResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}

type UploadResponse struct {
	Response
	Data *ImageInfo `json:"data"` // 上传的图片信息
}

type UserInfo struct {
	Id       int64  `json:"id"`                 // 用户ID
	Username string `json:"username"`           // 用户名
	Nickname string `json:"nickname"`           // 昵称
	Email    string `json:"email"`              // 邮箱
	Phone    string `json:"phone"`              // 手机号
	Avatar   string `json:"avatar"`             // 头像
	Gender   int32  `json:"gender,options=0|1"` // 性别 0:未知 1:男 2:女
	Age      int32  `json:"age,range=[0:120]"`  // 年龄
	Status   bool   `json:"status"`             // 状态 true:正常 false:禁用
	IsAdmin  bool   `json:"is_admin"`           // 角色 0:普通用户 1:管理员
	CreateAt int64  `json:"create_at"`          // 创建时间
	UpdateAt int64  `json:"update_at"`          // 更新时间
}

type UserList struct {
	Total int64      `json:"total"` // 总数
	List  []UserInfo `json:"list"`  // 用户列表
}
