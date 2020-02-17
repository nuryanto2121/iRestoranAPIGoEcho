package allmaster

import (
	"context"
	"yanto/irestora/models"
)

// Repository :
type Repository interface {
	GetAllMaster(ctx context.Context) (result models.AllMaster, err error)
}

// Usecase :
type Usecase interface {
	GetAllMaster(ctx context.Context) (result models.AllMaster, err error)
}
