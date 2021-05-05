package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"tasksvc/task-rest-api/internal/model"
)

func (h *Handler) createTask(c *gin.Context) {
	var input model.Task

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	taskId, err := h.services.CreateTask(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": taskId,
	})
}

func (h *Handler) getAllTasks(c *gin.Context) {
	tasks, err := h.services.GetAllTasks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) getTaskById(c *gin.Context) {
	idParam := c.Param("id")
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := h.services.GetTaskById(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) updateTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	taskId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "couldn't parse json request")
	}

	updated, err := h.services.UpdateTaskById(taskId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if updated {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": "updated successfully",
		})
	} else {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"result": "task not found",
		})
	}
}

func (h *Handler) deleteTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty id param")
		return
	}
	objectId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	deleted, err := h.services.DeleteTaskById(objectId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if deleted {
		c.JSON(http.StatusOK, map[string]interface{}{
			"result": "deleted successfully",
		})
	} else {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"result": "task not found",
		})
	}
}
