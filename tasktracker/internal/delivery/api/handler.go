package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"net/http"
)

type Handler struct {
	taskService services.TasksService
}

func NewHandler(service services.TasksService) *Handler {
	return &Handler{taskService: service}
}

func (h *Handler) taskAdd(c *gin.Context) {
	var err error
	req := new(TaskAddRequest)
	err = c.BindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ErrorResponse{
			true, fmt.Sprintf("%s", err.Error()),
		})
		return
	}

	var task *domain.Task
	task, err = h.taskService.AddTask(c.Request.Context(), req.Title, req.Title)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &ErrorResponse{
			true, fmt.Sprintf("%s", err.Error()),
		})
	}

	c.JSON(http.StatusOK, &TaskAddResponse{
		PublicId: task.PublicId,
	})
}
