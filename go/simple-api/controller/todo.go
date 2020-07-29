package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/Kudoas/sandbox/go/simple-api/repository"
	"github.com/gin-gonic/gin"
)

func FindTodos(c *gin.Context) {
	todos := repository.AllTodo()
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(c *gin.Context) {
	var todo model.InputTodo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo := model.Todo{Text: todo.Text, Status: todo.Status, CreateAt: time.Now()}
	log.Print(newTodo)
	repository.CreateTodo(&newTodo)
	log.Print(newTodo)
	c.JSON(http.StatusOK, gin.H{"data": newTodo})
}

func FindTodo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Print(err)
	}
	todo := repository.FindTodo(id)
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

// func UpdateTodo(c *gin.Context) {
// 	var todo model.Todo
// 	n := c.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	text := c.PostForm("text")
// 	status := c.PostForm("status")
// 	todo.ID = id
// 	todo.Text = text
// 	todo.Status = status
// 	repository.UpdateTodo(todo)
// 	c.JSON(http.StatusOK, gin.H{"todo": todo})
// }

// func DeleteTodo(c *gin.Context) {
// 	n := c.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	repository.DeleteTodo(id)
// 	c.JSON(http.StatusNotFound, nil)
// }
