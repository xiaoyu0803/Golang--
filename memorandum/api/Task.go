package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"todo_list/pkg/utils"
	"todo_list/service"
)

func CreataTask(c *gin.Context) {
	var createTask service.CreateTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.Println(err)
	}

}

func TaskListOne(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.ShowTask(c.Query("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.Println(err)
	}
}

func TaskListall(c *gin.Context) {
	var showTaskAll service.ShowTaskService
	if err := c.ShouldBind(&showTaskAll); err == nil {
		res := showTaskAll.ShowTaskAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.Println(err)
	}
}

func ModifyTask(c *gin.Context) {
	var modifyTask service.ModifyTask
	if err := c.ShouldBind(&modifyTask); err == nil {
		res := modifyTask.ModifyTask(c.Query("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.Println(err)
	}
}

func DeleteTask(c *gin.Context) {
	var modifyTask service.DeleteTask
	if err := c.ShouldBind(&modifyTask); err == nil {
		res := modifyTask.DeleteTask(c.Query("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.Println(err)
	}
}
