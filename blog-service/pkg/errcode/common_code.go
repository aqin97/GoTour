package errcode

var (
	Sucess                    = NewError(0, " 成功")
	ServerError               = NewError(10000, "服务器内部错误")
	InvalidParams             = NewError(10001, "参数错误")
	NotFound                  = NewError(10002, "找不到对应资源")
	UnauthorizedAuthNotExist  = NewError(10003, "鉴权失败， 找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10004, "鉴权失败， token错误")
	UnauthorizedTokenTimeOut  = NewError(10005, "鉴权失败， token超时")
	UnauthorizedTokenGenerate = NewError(10006, "鉴权失败， token生成失败")
	TooManyRequest            = NewError(10007, "请求过多")
)
