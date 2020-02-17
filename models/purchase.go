package models

import "time"

//TbldataPurchase : data post
type TbldataPurchase struct {
	ID                     int16                     `json:"id" db:"id"`
	ReferenceNo            string                    `json:"reference_no" db:"reference_no"`
	SupplierID             int16                     `json:"supplier_id" db:"supplier_id"`
	Date                   string                    `json:"date" db:"date"`
	Subtotal               float32                   `json:"subtotal" db:"subtotal"`
	Other                  float32                   `json:"other" db:"other"`
	GrandTotal             float32                   `json:"grand_total" db:"grand_total"`
	Paid                   float32                   `json:"paid" db:"paid"`
	Due                    float32                   `json:"due" db:"due"`
	Note                   *string                   `json:"note" db:"note"`
	UserID                 int16                     `json:"user_id" db:"user_id"`
	OutletID               int16                     `json:"outlet_id" db:"outlet_id"`
	DelStatus              *string                   `json:"del_status" db:"del_status"`
	TblPurchaseIngredients []*TblPurchaseIngredients `json:"TblPurchaseIngredients" db:"tbl_purchase_ingredients"`
}

//TblPurchase : nama model = nama table di DB
type TblPurchase struct {
	ID           int16     `json:"id" db:"id"`
	ReferenceNo  string    `json:"reference_no" db:"reference_no"`
	SupplierID   int16     `json:"supplier_id" db:"supplier_id"`
	Date         string    `json:"date" db:"date"`
	Subtotal     float32   `json:"subtotal" db:"subtotal"`
	Other        float32   `json:"other" db:"other"`
	GrandTotal   float32   `json:"grand_total" db:"grand_total"`
	Paid         float32   `json:"paid" db:"paid"`
	Due          float32   `json:"due" db:"due"`
	Note         *string   `json:"note" db:"note"`
	UserID       int16     `json:"user_id" db:"user_id"`
	OutletID     int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus    *string   `json:"del_status" db:"del_status"`
	IDFromOutlet int16     `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload bool      `json:"is_data_upload" db:"is_data_upload"`
	DateUpload   time.Time `json:"date_upload" db:"date_upload"`
}

// TblPurchaseIngredients :
type TblPurchaseIngredients struct {
	ID             int16     `json:"id" db:"id"`
	IngredientID   int16     `json:"ingredient_id" db:"ingredient_id"`
	UnitPrice      float32   `json:"unit_price" db:"unit_price"`
	QuantityAmount float32   `json:"quantity_amount" db:"quantity_amount"`
	Total          float32   `json:"total" db:"total"`
	PurchaseID     int16     `json:"purchase_id" db:"purchase_id"`
	OutletID       int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus      *string   `json:"del_status" db:"del_status"`
	IDFromOutlet   int16     `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload   bool      `json:"is_data_upload" db:"is_data_upload"`
	DateUpload     time.Time `json:"date_upload" db:"date_upload"`
}

// DataPurchase :
type DataPurchase struct {
	Purchase []*TbldataPurchase
}
