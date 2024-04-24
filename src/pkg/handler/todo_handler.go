package handler

import (
	"cassandra/src/cmd/utils"
	"cassandra/src/pkg/model"
	"fmt"
	"net/http"

	database "cassandra/src/pkg/repository/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoHandler interface {
	CreateTodo(*gin.Context)
	GetTodoById(*gin.Context)
	HealthCheck(*gin.Context)
}

type todoHandler struct {
	repo database.TodoRepository
}

// NewTodoHandler is a create new todoHandler struct
func NewTodoHandler(repo *database.TodoRepository) TodoHandler {
	return &todoHandler{repo: *repo}
}

func (t *todoHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func (t *todoHandler) CreateTodo(c *gin.Context) {
	var todo model.Todo

	c.BindJSON(&todo)

	todo.Id = uuid.New().String()
	data, err := t.repo.Save(todo)
	fmt.Printf("data err ", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequestError("insert operation failed!", err))
	}

	c.JSON(http.StatusCreated, gin.H{"todo": data})
}

func (t *todoHandler) GetTodoById(c *gin.Context) {

	id := c.Param("id")

	todo, err := t.repo.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.BadRequestError("todo not found", err))
	}

	c.JSON(http.StatusOK, gin.H{"todo": todo})
}
