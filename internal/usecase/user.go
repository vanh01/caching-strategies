package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/vanh01/caching-strategies/internal/model"
)

type userUsecase struct {
	userRepo UserRepo
}

func NewUserUsecase(userRepo UserRepo) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (a *userUsecase) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return a.userRepo.GetById(ctx, id)
}
