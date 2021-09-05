package service

import (
	"context"
	"go-rest-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryCreateResponse
	Update(ctx context.Context, request web.CategoryCreateUpdate) web.CategoryCreateResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryCreateResponse
	FindAll(ctx context.Context) []web.CategoryCreateResponse
}
