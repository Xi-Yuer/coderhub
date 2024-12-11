// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

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
	Data *List `json:"data"` // 学术导航列表
}

type HealthResp struct {
	Response
	Data bool `json:"data"` // 是否健康
}

type List struct {
	Total int64               `json:"total"` // 总数
	List  []AcademicNavigator `json:"list"`  // 学术导航列表
}

type PostAcademicNavigatorLikeReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type PostAcademicNavigatorLikeResp struct {
	Response
	Data bool `json:"data"` // 是否点赞成功
}

type Response struct {
	Code    int32  `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
}

type UpdateAcademicNavigatorLikeCountReq struct {
	Id int64 `path:"id"` // 学术导航 ID
}

type UpdateAcademicNavigatorLikeCountResp struct {
	Response
	Data bool `json:"data"` // 是否更新成功
}
