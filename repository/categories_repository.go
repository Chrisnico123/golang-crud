package repository

import (
	"context"
	"database/sql"
	"golang_api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx , categories domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx , categories domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx , categories domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoriesId int) (domain.Category , error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}