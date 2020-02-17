package httppurchase

import (
	"context"
	"net/http"
	"yanto/irestora/models"
	purchase "yanto/irestora/transaksi/purchase/domain"

	"github.com/labstack/echo"
)

// PurchaseHandler :
type PurchaseHandler struct {
	PurchaseUsecase purchase.Usecase
}

// NewPurchaseHandler :
func NewPurchaseHandler(e *echo.Echo, pU purchase.Usecase) {
	handler := &PurchaseHandler{
		PurchaseUsecase: pU,
	}
	e.GET("/purchase", handler.GetPurchase)
	e.POST("/purchase", handler.AddPurchase)
}

// GetPurchase :
func (p *PurchaseHandler) GetPurchase(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ListPurchase, err := p.PurchaseUsecase.GetPurchase(ctx)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, ListPurchase)
}

// AddPurchase :
func (p *PurchaseHandler) AddPurchase(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	// var dta = map[string]interface{}{}
	// jsonMap := make(map[string]interface{})
	dataPurchase := &models.DataPurchase{}
	// dataPurchase2 := &models.DataPurchase{}
	// // var dataDetailpurchase []models.TblPurchaseIngredients
	if err := e.Bind(dataPurchase); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	// err := json.NewDecoder(e.Request().Body).Decode(&jsonMap)
	// err = json.NewDecoder(e.Request().Body).Decode(&dta)
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, err.Error())
	// }
	// // json.Unmarshal([]byte(jsonMap), dataPurchase)
	// mapstructure.Decode(jsonMap, &dataPurchase)

	err := p.PurchaseUsecase.AddPurchase(ctx, dataPurchase)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, dataPurchase)
}
