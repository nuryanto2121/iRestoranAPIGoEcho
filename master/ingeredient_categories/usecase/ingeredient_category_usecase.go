package ingeredientcategoryusecase

import (
	"context"
	"time"
	ingeredientcategories "yanto/irestora/master/ingeredient_categories/domain"
	"yanto/irestora/models"
)

type ingeredientCategoryUsecase struct {
	ingeredientcategoryRepo ingeredientcategories.Repository
	contextTimeout          time.Duration
}

//NewIngeredientCategoryUsecase :
func NewIngeredientCategoryUsecase(a ingeredientcategories.Repository, timeout time.Duration) ingeredientcategories.Usecase {
	return &ingeredientCategoryUsecase{
		ingeredientcategoryRepo: a,
		contextTimeout:          timeout,
	}
}

func (a *ingeredientCategoryUsecase) GetIngeredientCategories(ctx context.Context) (result []models.IngredientCategories, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	listIngeredient, err := a.ingeredientcategoryRepo.GetIngeredientCategories(ctx)
	if err != nil {
		return nil, err

	}
	return listIngeredient, nil
}
