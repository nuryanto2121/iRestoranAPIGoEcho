package purchase

import (
	"context"
	"yanto/irestora/models"
)

// Repository :
type Repository interface {
	GetPurchase() (result []*models.TblPurchase, err error)
	AddPurchase(pu *models.TblPurchase) error
	AddPurchaseIngeredient(pu *models.TblPurchaseIngredients) error
}

// Usecase :
type Usecase interface {
	GetPurchase(ctx context.Context) (result []*models.TblPurchase, err error)
	// AddPurchase(ctx context.Context, pu models.DataPurchase) error
	AddPurchase(ctx context.Context, pu *models.DataPurchase) error
}
