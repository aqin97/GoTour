package errcode

var (
	Sucess                    = NewError(0, "成功")
	ServerError               = NewError(1001, "服务器内部错误")
	InvalidParams             = NewError(1002, "无效参数")
	NotFound                  = NewError(1004, "找不到")
	UnauthorizedAuthNotExist  = NewError(1005, "鉴权失败，找不到对应的AppKey和对应的AppSecret")
	UnauthorizedTokenError    = NewError(1006, "鉴权失败，token错误")
	UnauthorizedTokenTimeOut  = NewError(1007, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewError(1008, "鉴权失败，token生成失败")
	TooManyRequest            = NewError(1008, "请求过多")
)
