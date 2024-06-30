package logic

import (
	"fmt"
	"online-house-trading-platform/codes"
	"online-house-trading-platform/internal/dao"
	"online-house-trading-platform/pkg/model"
	"path/filepath"
)

// RegisterHandle 用于处理用户注册逻辑
func RegisterHandle(req model.RegisterRequest) *model.Error {
	if req.Username == "" || req.Password == "" || req.Email == "" {
		return &model.Error{StatusCode: codes.RegisterInvalidParam}
	}

	usernameExists, emailExists, err := dao.CheckUserExists(req.Username, req.Email)
	if err != nil {
		return &model.Error{StatusCode: codes.CheckUserExistsError}
	}
	if usernameExists {
		return &model.Error{StatusCode: codes.RegisterUsernameExists}
	}
	if emailExists {
		return &model.Error{StatusCode: codes.RegisterEmailExists}
	}

	//user := model.User{
	//	Username: req.Username,
	//	Password: EncryptPassword(req.Password),
	//	Email:    req.Email,
	//	Role:     req.Role,
	//}

	user := model.User{
		Username:     req.Username,
		Password:     EncryptPassword(req.Password),
		Email:        req.Email,
		Role:         req.Role,
		PhoneNumber:  req.PhoneNumber,
		Sex:          req.Sex,
		Province:     req.Province,
		City:         req.City,
		Identity:     req.Identity,
		QQNumber:     req.QQNumber,
		WechatNumber: req.WechatNumber,
	}

	err = dao.CreateUser(&user)
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterCreateUserError}
	}

	if req.Avatar == nil {
		return &model.Error{StatusCode: codes.RegisterInvalidParam, Message: "头像文件解析失败或者未携带头像文件.此时用户已经创建完成,无需再创建用户"}
	}
	fileName := generateRandomFileName()
	dst := fmt.Sprintf("./uploads/user/%d/%s%v", user.ID, fileName, filepath.Ext(req.Avatar.Filename))
	err = saveUploadedFile(req.Avatar, dst)
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterSaveAvatarError}
	}
	err = dao.CreateUserAvatar(&model.UserAvatar{UserID: user.ID, URL: dst})
	if err != nil {
		return &model.Error{StatusCode: codes.RegisterSaveAvatarError}
	}
	return nil
}
