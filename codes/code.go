package codes

type ResCode int64

const (
	Admin string = "admin"
	User  string = "user"
)

const (
	CodeSuccess                ResCode = 1000 + iota // 成功
	LoginInvalidParam                                // 请求参数错误
	LoginUserNotExist                                // 用户不存在
	LoginInvalidPassword                             // 密码错误
	LoginServerBusy                                  // 服务繁忙
	GetDBError                                       // 数据库错误
	GenerateJWTTokenError                            // 无法生成jwt token
	RequestWithoutTokenError                         // 请求未携带token，无权限访问
	InvalidTokenFormatError                          // token格式错误
	InvalidTokenError                                // token无效
	RegisterInvalidParam                             // 注册请求参数错误
	CheckUserExistsError                             // 检查用户是否存在错误
	RegisterUsernameExists                           // 用户名已存在
	RegisterEmailExists                              // 邮箱已存在
	RegisterCreateUserError                          // 创建用户失败
	ReserveInvalidParam                              // 预约请求参数错误
	ReserveError                                     // 预约失败
	GetUserIDError                                   // 获取用户ID错误
	UserIDTypeError                                  // 用户ID类型错误
	GetHouseListError                                // 获取房屋列表错误
	GetHouseInfoError                                // 获取某一个房屋信息错误
	HouseIDInvalid                                   // 房屋ID无效
	ReleaseBindDataError                             // 绑定数据错误
	CreateDirError                                   // 创建文件夹错误
	SaveFileError                                    // 保存文件错误
	CreateHouseError                                 // 创建房屋错误
	CreateHouseImageError                            // 创建房屋图片错误
	GetUserFavouritesError                           // 获取用户收藏错误
	GetUserProfileError                              // 获取用户信息错误
	BindDataError                                    // 绑定数据错误
	ModifyUserProfileError                           // 修改用户信息错误
	GetReserveInformationError                       // 获取预约信息错误
	UserIDNotMatch                                   // 用户ID不匹配
	DeleteHouseError                                 // 删除房屋错误
	UpdateHouseError                                 // 更新房屋错误
	RegisterSaveAvatarError                          // 保存用户头像错误
	DeleteUserAvatarError                            // 删除用户头像错误
	NoPermission                                     // 无权限
	GetAllUsersError                                 // 获取所有用户信息错误
	GetAllHousesError                                // 获取所有房屋信息错误
	OpenFileError                                    // 打开文件错误
	GetViewingRecordsError                           // 获取看房记录错误
	AddViewingRecordsError                           //增加看房记录错误
	RecordExists                                     //记录已存在
)

// codeMessageMap 用于存放code和message的映射关系
var codeMessageMap = map[ResCode]string{
	CodeSuccess:                "请求成功",
	LoginInvalidParam:          "登陆请求参数错误",
	LoginUserNotExist:          "登陆用户不存在",
	LoginInvalidPassword:       "登陆密码错误",
	LoginServerBusy:            "服务繁忙",
	GetDBError:                 "数据库错误",
	GenerateJWTTokenError:      "无法生成jwt token",
	RequestWithoutTokenError:   "请求未携带token，无权限访问",
	InvalidTokenFormatError:    "token格式错误",
	InvalidTokenError:          "token无效",
	RegisterInvalidParam:       "注册请求参数错误",
	CheckUserExistsError:       "检查用户是否存在错误",
	RegisterUsernameExists:     "注册时用户名已存在",
	RegisterEmailExists:        "注册时邮箱已存在",
	RegisterCreateUserError:    "创建用户失败",
	ReserveInvalidParam:        "预约请求参数错误",
	ReserveError:               "预约失败",
	GetUserIDError:             "获取用户ID错误",
	UserIDTypeError:            "用户ID类型错误",
	GetHouseListError:          "获取房屋列表错误",
	GetHouseInfoError:          "获取某一个房屋信息错误",
	HouseIDInvalid:             "房屋ID无效",
	ReleaseBindDataError:       "绑定数据错误",
	CreateDirError:             "创建文件夹错误",
	SaveFileError:              "保存文件错误",
	CreateHouseError:           "创建房屋错误",
	CreateHouseImageError:      "创建房屋图片错误",
	GetUserFavouritesError:     "获取用户收藏错误",
	GetUserProfileError:        "获取用户信息错误",
	BindDataError:              "绑定数据错误",
	ModifyUserProfileError:     "修改用户信息错误",
	GetReserveInformationError: "获取预约信息错误",
	UserIDNotMatch:             "Context中和url中的用户ID不匹配",
	DeleteHouseError:           "删除房屋错误",
	UpdateHouseError:           "更新房屋错误",
	RegisterSaveAvatarError:    "保存用户头像错误",
	DeleteUserAvatarError:      "删除用户头像错误",
	NoPermission:               "无权限",
	GetAllUsersError:           "获取所有用户信息错误",
	GetAllHousesError:          "获取所有房屋信息错误",
	OpenFileError:              "打开文件错误",
	GetViewingRecordsError:     "获取看房记录错误",
	AddViewingRecordsError:     "增加看房记录错误",
	RecordExists:               "记录已存在",
}

const (
	WebsocketSuccessMessage ResCode = 2000 + iota
	WebsocketSuccess
	WebsocketEnd
	WebsocketOnlineReply
	WebsocketOfflineReply
	WebsocketLimit
)

var websocketMessageMap = map[ResCode]string{
	WebsocketSuccessMessage: "消息发送成功",
	WebsocketSuccess:        "连接成功",
	WebsocketEnd:            "连接结束",
	WebsocketOnlineReply:    "对方在线",
	WebsocketOfflineReply:   "对方不在线",
	WebsocketLimit:          "连接数已达上限",
}

// Message 返回code对应的消息
func (c ResCode) Message() string {
	var msg string
	if c < 2000 {
		msg, _ = codeMessageMap[c]
	} else {
		msg, _ = websocketMessageMap[c]
	}
	return msg
}

func (c ResCode) Int() int {
	return int(c)
}
