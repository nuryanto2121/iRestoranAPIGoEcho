package ingeredientcategories

import "context"

import "yanto/irestora/models"

// Repository :
type Repository interface {
	GetIngeredientCategories(ctx context.Context) (result []models.IngredientCategories, err error)
}

// Usecase :
type Usecase interface {
	GetIngeredientCategories(ctx context.Context) (result []models.IngredientCategories, err error)
}
