package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/vanh01/caching-strategies/internal/model"
	"github.com/vanh01/caching-strategies/pkg/cache"
)

type userUsecase struct {
	userRepo   UserRepo
	cacheAside *cache.BaseCache
}

func NewUserUsecase(userRepo UserRepo, cache *cache.BaseCache) *userUsecase {
	return &userUsecase{
		userRepo:   userRepo,
		cacheAside: cache,
	}
}

func (a *userUsecase) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user *model.User
	err := a.cacheAside.GetObject(id.String(), &user)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		fmt.Printf("Get user from database by %s\n", id)
		user, err = a.userRepo.GetById(ctx, id)
		if err != nil {
			return nil, err
		}

		a.cacheAside.SetObject(id.String(), user, 5*60)
	}

	return user, nil
}

func (a *userUsecase) GetByIdWithoutCache(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return a.userRepo.GetById(ctx, id)
}
