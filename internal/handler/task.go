package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tranngoclam/go-msvc/internal/config"
	"net/http"
)

type Task struct {
	config config.Config
}

func (t *Task) Routes(ctx context.Context) http.Handler {
	e := gin.Default()

	e.POST("/tasks", t.HandleCreateTask())
	e.PATCH("/tasks/:id", t.HandleUpdateTask())
	e.DELETE("/tasks/:id", t.HandleDeleteTask())

	return e
}

func (t *Task) HandleCreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (t *Task) HandleUpdateTask() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (t *Task) HandleDeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
