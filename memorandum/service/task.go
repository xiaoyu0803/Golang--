package service

import (
	"log"
	"time"
	"todo_list/model"
	"todo_list/serializer"
)

type CreateTaskService struct {
	Title   string `json:"title,omitempty" form:"title"`
	Content string `json:"content,omitempty" form:"content"`
	Status  int    `json:"status,omitempty" form:"status"`
}

type ShowTaskService struct {
}

type DeleteTask struct {
}

type ModifyTask struct {
	ID      uint
	Title   string `json:"title,omitempty" form:"title" binding:"min=2,max=50"`
	Content string `json:"content,omitempty" form:"content" binding:"max=5000"`
	Status  string `json:"status,omitempty" form:"status"`
}

func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    "0",
		StartTime: time.Now().Unix(),
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		log.Println(err)
		return serializer.Response{
			Status: 400,
			Msg:    "",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
	}
}

func (service *ShowTaskService) ShowTask(id string) serializer.Response {
	var task model.Task
	err := model.DB.Where("id = ?", id).Find(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "",
		Error:  "",
	}
}

func (service *ShowTaskService) ShowTaskAll() serializer.Response {
	var task []model.Task
	err := model.DB.Find(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTaskAll(task),
		Msg:    "",
		Error:  "",
	}
}

func (receiver ModifyTask) ModifyTask(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Title = receiver.Title
	task.Content = receiver.Content
	task.Status = receiver.Status
	err := model.DB.Save(task).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   receiver,
		Msg:    "修改成功",
		Error:  "",
	}
}

func (d DeleteTask) DeleteTask(id string) serializer.Response {
	var task model.Task
	var cont int
	model.DB.Where("id=?", id).First(&task).Count(&cont)
	if cont < 1 {
		return serializer.Response{
			Status: 404,
			Data:   nil,
			Msg:    "未找相关信息",
			Error:  "",
		}
	}
	err := model.DB.Delete(&task).Error
	if err != nil {
		return serializer.Response{
			400,
			"",
			"数据库操作错误",
			err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   task.ID,
		Msg:    "删除成功",
		Error:  "",
	}
}
