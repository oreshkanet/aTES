package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"net/http"
)

type Api struct {
	srv *http.Server

	handler *Handler
}

func NewApi(srv *http.Server, taskService services.TasksService) *Api {
	return &Api{
		srv: srv,
		handler: &Handler{
			taskService: taskService,
		},
	}
}

func (a *Api) Run() error {
	// Создаём новый роутер
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	a.srv.Handler = router

	// Роутим эндпойнт авторизации
	taskRouter := router.Group("/task")
	taskRouter.PUT("/", a.handler.taskAdd)
	//TODO: Добавить все роуты работы с задачами

	return a.srv.ListenAndServe()
}

func (a *Api) Stop() error {
	return a.srv.Close()
}
