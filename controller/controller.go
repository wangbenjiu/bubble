package controller

import (
	"bubble/dao"
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandlers(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	dao.DB.Create(&todo)
}

func GetTodoList(c *gin.Context) {
	var todolist []*models.Todo
	dao.DB.Find(&todolist)
	c.JSON(http.StatusOK, &todolist)
}

func UpdateATodo(c *gin.Context) {

}

func DeleteATodo(c *gin.Context) {

}
