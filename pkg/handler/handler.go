package handler

import (
	"github.com/gin-gonic/gin"
	"restful-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			tasks := lists.Group(":id/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/", h.getAllTasks)

				subtasks := tasks.Group(":task_id/subtasks")
				{
					subtasks.POST("/", h.createSubtask)
					subtasks.GET("/", h.getAllSubtasks)
				}
			}
		}
		tasks := api.Group("tasks")
		{
			tasks.GET("/:id", h.getTaskById)
			tasks.PUT("/:id", h.updateTask)
			tasks.DELETE("/:id", h.deleteTask)
		}
		subtasks := api.Group("subtasks")
		{
			subtasks.GET("/:id", h.getSubtaskById)
			subtasks.PUT("/:id", h.updateSubtask)
			subtasks.DELETE("/:id", h.deleteSubtask)
		}
	}
	return router
}
