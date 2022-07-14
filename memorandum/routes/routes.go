package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"todo_list/api"
	"todo_list/middleware"
)

func NewRoute() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/task")
		authed.Use(middleware.JWT())
		{
			authed.POST("/create", api.CreataTask)
			authed.GET("/listone", api.TaskListOne)
			authed.GET("/listall", api.TaskListall)
			authed.PUT("/modify", api.ModifyTask)
			authed.DELETE("/deletetask", api.DeleteTask)
		}
	}
	return r
}
