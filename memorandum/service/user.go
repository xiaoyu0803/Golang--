package service

import (
	"log"
	"todo_list/model"
	"todo_list/pkg/utils"
	"todo_list/serializer"

	"github.com/jinzhu/gorm"
)

type UsersService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=5"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

func (service *UsersService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).
		Where("user_name", service.UserName).Find(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经被注册",
		}
	}
	user.UserName = service.UserName
	err := user.SetPassword(service.Password)
	if err != nil {
		log.Println(err)
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库操作错误",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   user.UserName,
		Msg:    "用户注册成功",
	}
}

func (service *UsersService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//token，为了其他功能需要身份验证给前端存储
	//创建备忘录需要tocken
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		log.Println(err)
		return serializer.Response{
			Status: 500,
			Msg:    "token错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
