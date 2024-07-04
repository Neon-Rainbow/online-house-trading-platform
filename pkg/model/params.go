package model

import (
	"mime/multipart"
	"online-house-trading-platform/codes"
)

// LoginRequest 用于处理用户登录请求
type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// LoginResponse 用于处理用户登录请求
type LoginResponse struct {
	AccessToken  string `json:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
	UserID       uint   `json:"user_id" form:"user_id"`
	Username     string `json:"username" form:"username"`
	Role         string `json:"role" form:"role"`
}

// RegisterRequest 用于处理用户注册请求
type RegisterRequest struct {
	Username     string                `json:"username" form:"username"`
	Password     string                `json:"password" form:"password"`
	Email        string                `json:"email" form:"email"`
	Role         string                `json:"role" form:"role"`
	PhoneNumber  string                `json:"phone_number" form:"phone_number"` // PhoneNumber 用户手机号
	Sex          string                `json:"sex" form:"sex"`
	Province     string                `json:"province" form:"province" gorm:"type:varchar(255)"`
	City         string                `json:"city" form:"city" gorm:"type:varchar(255)"`
	Identity     string                `json:"identity" form:"identity" gorm:"type:varchar(255)"`
	QQNumber     string                `json:"qq_number" form:"qq_number" gorm:"type:varchar(255)"`
	WechatNumber string                `json:"wechat_number" form:"wechat_number" gorm:"type:varchar(255)"`
	Avatar       *multipart.FileHeader `json:"avatar" form:"avatar" `
}

func (r *RegisterRequest) ConvertUserModelWithoutAvatar() *User {
	return &User{
		Username:     r.Username,
		Password:     r.Password,
		Email:        r.Email,
		Role:         r.Role,
		PhoneNumber:  r.PhoneNumber,
		Sex:          r.Sex,
		Province:     r.Province,
		City:         r.City,
		Identity:     r.Identity,
		QQNumber:     r.QQNumber,
		WechatNumber: r.WechatNumber,
	}
}

type AdminRegisterRequest struct {
	Username     string                `json:"username" form:"username"`
	Password     string                `json:"password" form:"password"`
	Email        string                `json:"email" form:"email"`
	Role         string                `json:"role" form:"role"`
	PhoneNumber  string                `json:"phone_number" form:"phone_number"` // PhoneNumber 用户手机号
	Sex          string                `json:"sex" form:"sex"`
	Province     string                `json:"province" form:"province" gorm:"type:varchar(255)"`
	City         string                `json:"city" form:"city" gorm:"type:varchar(255)"`
	Identity     string                `json:"identity" form:"identity" gorm:"type:varchar(255)"`
	QQNumber     string                `json:"qq_number" form:"qq_number" gorm:"type:varchar(255)"`
	WechatNumber string                `json:"wechat_number" form:"wechat_number" gorm:"type:varchar(255)"`
	Avatar       *multipart.FileHeader `json:"avatar" form:"avatar" `
	AdminSecret  string                `json:"admin_secret" form:"admin_secret"`
}

// HouseRequest 是从前端接收房源数据的结构体。
type HouseRequest struct {
	Owner            string                  `json:"owner" form:"owner"`                         // 房屋所有者名称
	OwnerID          uint                    `json:"owner_id" form:"owner_id"`                   // 房屋所有者ID
	Title            string                  `json:"title" form:"title"`                         // 房屋标题
	Description      string                  `json:"description" form:"description"`             // 房屋描述
	Price            float64                 `json:"price" form:"price"`                         // 房屋价格
	Address          string                  `json:"address" form:"address"`                     // 房屋地址
	HouseOrientation string                  `json:"house_orientation" form:"house_orientation"` // HouseOrientation 房屋朝向
	Layout           string                  `json:"layout" form:"layout"`                       // Layout 房屋户型
	Area             float64                 `json:"area" form:"area"`                           // Area 房屋面积
	Floor            string                  `json:"floor" form:"floor"`                         // Floor 房屋楼层
	RentPrice        float64                 `json:"rent_price" form:"rent_price"`               // RentPrice 房屋租金
	Type             string                  `json:"type" form:"type"`                           // Type 房屋类型
	PostCode         string                  `json:"post_code" form:"post_code"`                 // PostCode 房屋邮编
	Images           []*multipart.FileHeader `json:"images" form:"images"`                       // 房屋相关的图片文件列表
}

// ConvertToHouseModel 将 HouseRequest 转换为 House 模型
func (req *HouseRequest) ConvertToHouseModel() *House {
	house := &House{
		Owner:            req.Owner,
		OwnerID:          req.OwnerID,
		Title:            req.Title,
		Description:      req.Description,
		Price:            req.Price,
		Address:          req.Address,
		HouseOrientation: req.HouseOrientation,
		Layout:           req.Layout,
		Area:             req.Area,
		Floor:            req.Floor,
		RentPrice:        req.RentPrice,
		Type:             req.Type,
		PostCode:         req.PostCode,
	}
	return house
}

type HouseUpdateRequest struct {
	HouseID          uint                    `json:"house_id" form:"house_id"`                   // 房屋ID
	Owner            string                  `json:"owner" form:"owner"`                         // 房屋所有者名称
	OwnerID          uint                    `json:"owner_id" form:"owner_id"`                   // 房屋所有者ID
	Title            string                  `json:"title" form:"title"`                         // 房屋标题
	Description      string                  `json:"description" form:"description"`             // 房屋描述
	Price            float64                 `json:"price" form:"price"`                         // 房屋价格
	Address          string                  `json:"address" form:"address"`                     // 房屋地址
	HouseOrientation string                  `json:"house_orientation" form:"house_orientation"` // HouseOrientation 房屋朝向
	Layout           string                  `json:"layout" form:"layout"`                       // Layout 房屋户型
	Area             float64                 `json:"area" form:"area"`                           // Area 房屋面积
	Floor            string                  `json:"floor" form:"floor"`                         // Floor 房屋楼层
	RentPrice        float64                 `json:"rent_price" form:"rent_price"`               // RentPrice 房屋租金
	Type             string                  `json:"type" form:"type"`                           // Type 房屋类型
	PostCode         string                  `json:"post_code" form:"post_code"`                 // PostCode 房屋邮编
	Images           []*multipart.FileHeader `json:"images" form:"images"`                       // 房屋相关的图片文件列表
}

type UserReq struct {
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	Email        string `json:"email" form:"email"`
	Role         string `json:"role" form:"role"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"` // PhoneNumber 用户手机号
	Sex          string `json:"sex" form:"sex"`
	Province     string `json:"province" form:"province" gorm:"type:varchar(255)"`
	City         string `json:"city" form:"city" gorm:"type:varchar(255)"`
	Identity     string `json:"identity" form:"identity" gorm:"type:varchar(255)"`
	QQNumber     string `json:"qq_number" form:"qq_number" gorm:"type:varchar(255)"`
	WechatNumber string `json:"wechat_number" form:"wechat_number" gorm:"type:varchar(255)"`
}

type UserAvatarReq struct {
	UserID uint                  `json:"user_id" form:"user_id"`
	Avatar *multipart.FileHeader `json:"avatar" form:"avatar"`
}

// Error 用于处理错误信息
type Error struct {
	StatusCode codes.ResCode
	Message    string
}
