package errcode

var (
	ErrorGetTagListFail = NewError(20001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(20002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(20003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(20004, "删除标签失败")
	ErrorCountTagFail   = NewError(20005, "统计标签失败")
)
