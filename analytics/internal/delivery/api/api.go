package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/analytics/internal/service"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"net/http"
)

type Api struct {
	srv     *http.Server
	auth    authorizer.AuthToken
	handler *Handler
}

type Config struct {
	Srv      *http.Server
	Auth     authorizer.AuthToken
	Analytic service.Analytic
}

func NewApi(config *Config) *Api {
	return &Api{
		srv:  config.Srv,
		auth: config.Auth,
		handler: &Handler{
			analytic: config.Analytic,
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
	taskRouter := router.Group("/accounting")
	taskRouter.GET("/balance", a.handler.getBalance)
	//TODO: Добавить все роуты работы с аккаунтингом

	return a.srv.ListenAndServe()
}

func (a *Api) Stop() error {
	return a.srv.Close()
}
