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
	userRepo UserRepo
	redis    *cache.BaseCache
}

func NewUserUsecase(userRepo UserRepo, cache *cache.BaseCache) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		redis:    cache,
	}
}

func (a *userUsecase) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var user *model.User
	err := a.redis.GetObject(id.String(), &user)
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if err == redis.Nil {
		fmt.Println("Get data from database")
		user, err = a.userRepo.GetById(ctx, id)
		if err != nil {
			return nil, err
		}

		a.redis.SetObject(id.String(), user, 5*60)
	}

	return user, nil
}
