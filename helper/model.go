package helper

import (
	"go-restful-api/model/domain"
	"go-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var response []web.CategoryResponse
	for _, category := range categories {
		response = append(response, ToCategoryResponse(category))
	}

	return response
}
