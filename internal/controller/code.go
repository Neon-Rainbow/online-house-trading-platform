package controller

type ResCode int64

const (
	CodeSuccess           ResCode = 1000 + iota // 成功
	LoginInvalidParam                           // 请求参数错误
	LoginUserNotExist                           // 用户不存在
	LoginInvalidPassword                        // 密码错误
	LoginServerBusy                             // 服务繁忙
	GetDBError                                  // 数据库错误
	GenerateJWTTokenError                       //无法生成jwt token
)

// codeMessageMap 用于存放code和message的映射关系
var codeMessageMap = map[ResCode]string{
	CodeSuccess:           "请求成功",
	LoginInvalidParam:     "登陆请求参数错误",
	LoginUserNotExist:     "登陆用户不存在",
	LoginInvalidPassword:  "登陆密码错误",
	LoginServerBusy:       "服务繁忙",
	GetDBError:            "数据库错误",
	GenerateJWTTokenError: "无法生成jwt token",
}

// Message 返回code对应的消息
func (c ResCode) Message() string {
	msg, ok := codeMessageMap[c]
	if !ok {
		return codeMessageMap[LoginServerBusy]
	}
	return msg
}
