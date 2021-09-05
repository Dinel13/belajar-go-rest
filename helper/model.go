package helper

import (
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryCreateResponse {
	return web.CategoryCreateResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryCreateResponse {
	var categoryResponses []web.CategoryCreateResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
