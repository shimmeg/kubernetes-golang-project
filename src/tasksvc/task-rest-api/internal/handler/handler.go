package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasksvc/task-rest-api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	health := router.Group("/health")
	{
		health.GET("", h.healthCheck)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.POST("", h.createTask)
			tasks.GET("", h.getAllTasks)
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
	}
	return router
}

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"health": "ok"})
}
