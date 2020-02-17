package httpingeredientcategory

import (
	"context"
	"net/http"
	ingeredientcategories "yanto/irestora/master/ingeredient_categories/domain"
	"yanto/irestora/models"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// ResponseError :
type ResponseError struct {
	Message string `json:"message"`
}

// IngeredientCategoriesHandler :
type IngeredientCategoriesHandler struct {
	ICusecase ingeredientcategories.Usecase
}

// NewIngeredientCategoriesHandler :
func NewIngeredientCategoriesHandler(e *echo.Echo, us ingeredientcategories.Usecase) {
	handler := &IngeredientCategoriesHandler{
		ICusecase: us,
	}

	e.GET("/master/ingeredient", handler.GetIngeredientCategories)
}

// GetIngeredientCategories :
func (I *IngeredientCategoriesHandler) GetIngeredientCategories(c echo.Context) error {
	// var ingeredient []models.IngeedientCategories

	// err := c.Bind(&ingeredient)
	// if err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }

	// if ok, err := isRequestValid(&ingeredient); !ok {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listarr, err := I.ICusecase.GetIngeredientCategories(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, listarr)
}

func isRequestValid(m *models.IngredientCategories) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
