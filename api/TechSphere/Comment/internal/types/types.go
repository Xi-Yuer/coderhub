// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Comment struct {
	Id        int64          `json:"id"`         // 评论ID
	ArticleId int64          `json:"article_id"` // 文章ID
	Content   string         `json:"content"`    // 评论内容
	ParentId  int64          `json:"parent_id"`  // 父评论ID
	UserId    int64          `json:"user_id"`    // 评论者ID
	CreatedAt int64          `json:"created_at"` // 创建时间
	UpdatedAt int64          `json:"updated_at"` // 更新时间
	Replies   []*Comment     `json:"replies"`    // 子评论列表
	LikeCount int32          `json:"like_count"` // 点赞数
	Images    []CommentImage `json:"images"`     // 评论图片列表
}

type CommentImage struct {
	ImageId      string `json:"image_id"`      // 图片ID
	Url          string `json:"url"`           // 图片URL
	ThumbnailUrl string `json:"thumbnail_url"` // 缩略图URL
}

type CreateCommentReq struct {
	ArticleId  int64    `json:"article_id"`   // 文章ID
	Content    string   `json:"content"`      // 评论内容
	ParentId   int64    `json:"parent_id"`    // 父评论ID（可选）
	ReplyToUID int64    `json:"reply_to_uid"` // 回复的目标评论ID（可选）
	ImageIds   []string `json:"image_ids"`    // 图片ID列表
}

type CreateCommentResp struct {
	Response
	Data Comment `json:"data"` // 创建的评论
}

type DeleteCommentReq struct {
	CommentId int64 `path:"comment_id"` // 评论ID
}

type DeleteCommentResp struct {
	Response
	Data bool `json:"data"` // 删除是否成功
}

type GetCommentReq struct {
	CommentId int64 `path:"comment_id"` // 评论ID
}

type GetCommentResp struct {
	Response
	Data Comment `json:"data"` // 评论详情
}

type GetCommentsReq struct {
	ArticleId int64 `path:"article_id"` // 文章ID
	Page      int32 `form:"page"`       // 页码
	PageSize  int32 `form:"page_size"`  // 每页数量
}

type GetCommentsResp struct {
	Response
	Data List `json:"data"` // 评论列表
}

type HealthResp struct {
	Response
}

type List struct {
	List  []*Comment `json:"comments"` // 评论列表
	Total int32      `json:"total"`    // 总评论数
}

type Response struct {
	Code    int32  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
}
