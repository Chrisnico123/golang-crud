package service

import (
	"context"
	"database/sql"
	"golang_api/helper"
	"golang_api/model/domain"
	"golang_api/model/web"
	"golang_api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository , DB *sql.DB , validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context , request web.CategoryCreateRequest) web.CategoryResponse {
	errValidate :=	service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category {
		Name : request.Name,
	}

	category = service.CategoryRepository.Create(ctx ,tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceImpl) Update(ctx context.Context , request web.CategoryUpdateRequest) web.CategoryResponse {
	errValidate :=	service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category , err := service.CategoryRepository.FindById(ctx , tx , request.Id)
	helper.PanicIfError(err)

	category = service.CategoryRepository.Update(ctx ,tx, category)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceImpl) Delete(ctx context.Context , categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx , tx , categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx,tx ,category)
}
func (service *CategoryServiceImpl) FindById(ctx context.Context , categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx , tx , categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx , tx)

	return helper.ToCategoryResponses(categories)
}



