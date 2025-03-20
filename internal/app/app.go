package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"github.com/vanh01/caching-strategies/config"
	"github.com/vanh01/caching-strategies/internal/controller"
	"github.com/vanh01/caching-strategies/internal/repo"
	"github.com/vanh01/caching-strategies/internal/usecase"
	"github.com/vanh01/caching-strategies/pkg/cache"
)

func Run() {
	e := echo.New()

	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: false,
	}))

	redisConfig := config.Instance.RedisConfig
	client, err := cache.ConnectToRedis(fmt.Sprintf("redis://:%s@%s:%d/%d", redisConfig.Password, redisConfig.Host, redisConfig.Port, redisConfig.DB))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	pgConfig := config.Instance.PostgreConfig
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.DBName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Err to ping to database: %s\n", err.Error())
		return
	}

	cache := cache.NewBaseCache(client)
	userRepo := repo.NewUserRepo(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	controller.New(e, controller.UsecaseParam{
		UserUsecase: userUsecase,
		BaseCache:   cache,
	})

	address := fmt.Sprintf(":%d", config.Instance.Port)

	log.Fatal(e.Start(address))

	log.Println("Server exited!")
}
