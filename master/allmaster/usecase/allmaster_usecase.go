package allmasterusecase

import (
	"context"
	"time"
	allmaster "yanto/irestora/master/allmaster/domain"
	"yanto/irestora/models"
)

// allMasterUsecase :
type allMasterUsecase struct {
	allMasterRepo  allmaster.Repository
	contextTimeout time.Duration
}

// NewallMasterUsecase :
func NewallMasterUsecase(a allmaster.Repository, timeout time.Duration) allmaster.Usecase {
	return &allMasterUsecase{
		allMasterRepo:  a,
		contextTimeout: timeout,
	}
}

func (a *allMasterUsecase) GetAllMaster(ctx context.Context) (result models.AllMaster, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	listAllMaster, err := a.allMasterRepo.GetAllMaster(ctx)
	if err != nil {

		return models.AllMaster{}, err
	}

	return listAllMaster, nil
}
