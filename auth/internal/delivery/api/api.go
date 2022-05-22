package api

import (
	"github.com/gin-gonic/gin"
	"github.com/oreshkanet/aTES/auth/internal/service"
	"github.com/oreshkanet/aTES/packages/pkg/authorizer"
	"net/http"
)

type Api struct {
	srv  *http.Server
	auth authorizer.AuthToken
	h    *handler
}

type Config struct {
	Srv         *http.Server
	Auth        authorizer.AuthToken
	AuthService service.AuthService
}

func NewApi(config *Config) *Api {
	return &Api{
		srv:  config.Srv,
		auth: config.Auth,
		h:    NewHandler(config.AuthService),
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
	authRouter := router.Group("/auth")
	authRouter.POST("/sign_up", a.h.signUp)
	authRouter.POST("/sign_in", a.h.signIn)

	// Роутим эндпойнт управление профилем пользователя
	userRouter := router.Group("/user")
	userRouter.Use(a.UserMiddleware())
	userRouter.POST("/change_role", a.h.userChangeRole)
	userRouter.PUT("/", a.h.userUpdateProfile)

	return a.srv.ListenAndServe()
}

func (a *Api) Stop() error {
	return a.srv.Close()
}
