package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", CreateTodo)
		v1Group.GET("/todo", GetTodoList)
		v1Group.PUT("/todo/:id", UpdateATodo)
		v1Group.DELETE("/todo/:id", DeleteATodo)
	}

	return r
}

func IndexHandlers(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
