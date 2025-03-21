package controller

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/vanh01/caching-strategies/internal/usecase"
	"github.com/vanh01/caching-strategies/pkg/cache"
)

type UsecaseParam struct {
	UserUsecase usecase.UserUsecase
	BaseCache   *cache.BaseCache
}

func New(e *echo.Echo, params UsecaseParam) {
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	userRouter := &UserRouter{
		UserUsecase: params.UserUsecase,
	}

	g := e.Group("/api/v1")
	{
		g.GET("", func(c echo.Context) error {
			return c.String(http.StatusOK, "Application is running")
		})

		g.GET("/user/me", userRouter.GetMe)
	}
}

type UserRouter struct {
	UserUsecase usecase.UserUsecase
}

func (u *UserRouter) GetMe(c echo.Context) error {
	id := c.Request().Header.Get("userId")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := u.UserUsecase.GetById(context.Background(), userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
