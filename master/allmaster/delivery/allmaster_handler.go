package httpallmaster

import (
	"context"
	"net/http"
	allmaster "yanto/irestora/master/allmaster/domain"

	"github.com/labstack/echo"
)

// AllMasterHandler :
type AllMasterHandler struct {
	AllMasterUsecase allmaster.Usecase
}

// NewAllMasterHandler :
func NewAllMasterHandler(e *echo.Echo, amu allmaster.Usecase) {
	handler := &AllMasterHandler{
		AllMasterUsecase: amu,
	}

	e.GET("/master", handler.GetAllMaster)
}

//GetAllMaster :
func (A *AllMasterHandler) GetAllMaster(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	listarr, err := A.AllMasterUsecase.GetAllMaster(ctx) //I.ICusecase.GetIngeredientCategories(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, listarr)
}
