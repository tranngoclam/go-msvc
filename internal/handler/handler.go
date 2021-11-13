package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tranngoclam/go-msvc/internal/database"
	"github.com/tranngoclam/go-msvc/internal/model"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

var (
	healthzRateLimiter = rate.NewLimiter(rate.Every(time.Second), 1)
)

type Handler struct {
	mux http.Handler
	db  *database.Task
}

func NewHandler(db *database.Task) *Handler {
	h := &Handler{}

	m := gin.Default()

	m.Use(gin.Recovery())

	m.POST("/tasks", h.HandleCreateTask())
	m.GET("/tasks/:id", h.HandleGetTask())

	m.GET("/ping", h.HandlePing())
	m.GET("/healthz", h.HandleHealthz(db))

	h.mux = m
	h.db = db

	return h
}

func (h *Handler) HandleCreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var t model.Task
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if err := h.db.Put(c, &t); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf(`{"id": %d}`, t.ID))
	}
}

func (h *Handler) HandleGetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		task, err := h.db.Get(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, task)
	}
}

func (h *Handler) HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, []byte(`{"status": "ok"}`))
	}
}

func (h *Handler) HandleHealthz(db *database.Task) gin.HandlerFunc {
	return func(c *gin.Context) {
		if healthzRateLimiter.Allow() {
			err := db.Ping(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		c.JSON(http.StatusOK, []byte(`{"status": "ok"}`))
	}
}
