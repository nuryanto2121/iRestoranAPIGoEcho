package models

import "time"

// TblInventoryAdjustment :
type TblInventoryAdjustment struct {
	ID          int16     `json:"id" db:"id"`
	ReferenceNo *string   `json:"reference_no" db:"reference_no"`
	Date        time.Time `json:"date" db:"date"`
	Note        *string   `json:"note" db:"note"`
	EmployeeID  *int16    `json:"employee_id" db:"employee_id"`
	UserID      *int16    `json:"user_id" db:"user_id"`
	OutletID    *int16    `json:"outlet_id" db:"outlet_id"`
	DelStatus   *string   `json:"del_status" db:"del_status"`
}

// TblInventoryAdjustmentIngredients :
type TblInventoryAdjustmentIngredients struct {
	ID                    int16    `json:"id" db:"id"`
	IngredientID          *int16   `json:"ingredient_id" db:"ingredient_id"`
	ConsumptionAmount     *float32 `json:"consumption_amount" db:"consumption_amount"`
	InventoryAdjustmentID *int16   `json:"inventory_adjustment_id" db:"inventory_adjustment_id"`
	ConsumptionStatus     *string  `json:"consumption_status" db:"consumption_status"`
	OutletID              *int16   `json:"outlet_id" db:"outlet_id"`
	DelStatus             *string  `json:"del_status" db:"del_status"`
}
