package repopurchase

import "github.com/jinzhu/gorm"

import purchase "yanto/irestora/transaksi/purchase/domain"

import "yanto/irestora/models"

type mysqlPurchaseRepository struct {
	Conn *gorm.DB
}

// NewMysqlPurchaseRepository :
func NewMysqlPurchaseRepository(Conn *gorm.DB) purchase.Repository {
	return &mysqlPurchaseRepository{Conn}
}

// GetPurchase :
func (m *mysqlPurchaseRepository) GetPurchase() (result []*models.TblPurchase, err error) {
	// var
	err = m.Conn.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (m *mysqlPurchaseRepository) AddPurchase(purchase *models.TblPurchase) (err error) {
	err = m.Conn.Create(&purchase).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *mysqlPurchaseRepository) AddPurchaseIngeredient(pu *models.TblPurchaseIngredients) (err error) {
	err = m.Conn.Create(&pu).Error
	if err != nil {
		return err
	}
	return nil
}
