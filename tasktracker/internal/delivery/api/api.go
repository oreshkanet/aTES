package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/tasktracker/internal/services"
	"github.com/oreshkanet/aTES/tasktracker/pkg/authorizer"
	"net/http"
)

type Api struct {
	srv     *http.Server
	auth    authorizer.AuthToken
	handler *Handler
}

type Config struct {
	Srv         *http.Server
	Auth        authorizer.AuthToken
	TaskService services.TasksService
}

func NewApi(config *Config) *Api {
	return &Api{
		srv:  config.Srv,
		auth: config.Auth,
		handler: &Handler{
			taskService: config.TaskService,
		},
	}
}

func (a *Api) Run() error {
	// Создаём новый роутер
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
		a.AuthMiddleware(),
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
