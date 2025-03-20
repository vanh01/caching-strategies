package repo

import (
	"context"
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/vanh01/caching-strategies/internal/model"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (a *userRepo) GetById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	rows, err := squirrel.
		Select("*").
		From("user").
		Limit(1).
		Where(squirrel.Eq{"id": id}, id).
		RunWith(a.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
	}

	return &user, nil
}
