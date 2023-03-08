package controller

import (
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
	//dao.DB.Create(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)

	}

}

func GetTodoList(c *gin.Context) {
	var todolist []*models.Todo
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
	//dao.DB.Find(&todolist)
	//c.JSON(http.StatusOK, &todolist)
}

func UpdateATodo(c *gin.Context) {
	Id := c.Query("id")
}

func DeleteATodo(c *gin.Context) {

}
