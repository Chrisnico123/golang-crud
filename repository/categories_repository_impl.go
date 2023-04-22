package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_api/helper"
	"golang_api/model/domain"
)

type CategoryRespositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRespositoryImpl{}
}

func (repository *CategoryRespositoryImpl) Create(ctx context.Context , tx *sql.Tx , categories domain.Category) domain.Category {
	sql := "INSERT INTO categories(name) VALUES (?)"
	result , err := tx.ExecContext(ctx ,sql , categories.Name)
	helper.PanicIfError(err)

	id , err := result.LastInsertId()
	helper.PanicIfError(err)
	categories.Id = int(id)
	return categories
}

func (repository *CategoryRespositoryImpl) Update(ctx context.Context , tx *sql.Tx , categories domain.Category) domain.Category {
	sql := "UPDATE categories SET name = ? WHERE id = ?"
	_ , err := tx.ExecContext(ctx , sql , categories.Name , categories.Id)
	helper.PanicIfError(err)

	return categories
}

func (repository *CategoryRespositoryImpl) Delete(ctx context.Context , tx *sql.Tx , categories domain.Category) {
	sql := "DELETE FROM categories WHERE id = ?"
	_ , err := tx.ExecContext(ctx , sql ,categories.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRespositoryImpl) FindById(ctx context.Context, tx *sql.Tx , categoriesId int) (domain.Category , error ){
	sql := "SELECT id , name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx , sql , categoriesId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id , &category.Name)
		helper.PanicIfError(err)
		return category , nil
	} else {
		return category ,errors.New("category is not found")
	}


}

func (repository *CategoryRespositoryImpl) FindAll(ctx context.Context , tx *sql.Tx) []domain.Category {
	sql := "SELECT id , name FROM categories"
	rows , err := tx.QueryContext(ctx , sql)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id , &category.Name)
		helper.PanicIfError(err)
		categories = append(categories , category)
	}
	return categories

}