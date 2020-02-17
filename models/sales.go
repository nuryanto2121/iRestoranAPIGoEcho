package models

import "time"

// TblSales :
type TblSales struct {
	ID                      int16     `json:"id" db:"id"`
	CustomerID              string    `jsno:"customer_id" db:"customer_id"`
	SaleNo                  string    `jsno:"sale_no" db:"sale_no"`
	TotalItems              int16     `jsno:"total_items" db:"total_items"`
	SubTotal                float32   `jsno:"sub_total" db:"sub_total"`
	PaidAmount              float64   `jsno:"paid_amount" db:"paid_amount"`
	DueAmount               float32   `jsno:"due_amount" db:"due_amount"`
	Disc                    string    `jsno:"disc" db:"disc"`
	DiscActual              float32   `jsno:"disc_actual" db:"disc_actual"`
	Vat                     float32   `jsno:"vat" db:"vat"`
	TotalPayable            float32   `jsno:"total_payable" db:"total_payable"`
	PaymentMethodID         int16     `jsno:"payment_method_id" db:"payment_method_id"`
	CloseTime               time.Time `jsno:"close_time" db:"close_time"`
	TableID                 int16     `jsno:"table_id" db:"table_id"`
	TotalItemDiscountAmount float32   `jsno:"total_item_discount_amount" db:"total_item_discount_amount"`
	SubTotalWithDiscount    float32   `jsno:"sub_total_with_discount" db:"sub_total_with_discount"`
	SubTotalDiscountAmount  float32   `jsno:"sub_total_discount_amount" db:"sub_total_discount_amount"`
	TotalDiscountAmount     float32   `jsno:"total_discount_amount" db:"total_discount_amount"`
	DeliveryCharge          float32   `jsno:"delivery_charge" db:"delivery_charge"`
	SubTotalDiscountValue   string    `jsno:"sub_total_discount_value" db:"sub_total_discount_value"`
	SubTotalDiscountType    string    `jsno:"sub_total_discount_type" db:"sub_total_discount_type"`
	SaleDate                string    `jsno:"sale_date" db:"sale_date"`
	DateTime                time.Time `jsno:"date_time" db:"date_time"`
	OrderTime               time.Time `jsno:"order_time" db:"order_time"`
	CookingStartTime        time.Time `jsno:"cooking_start_time" db:"cooking_start_time"`
	CookingDoneTime         time.Time `jsno:"cooking_done_time" db:"cooking_done_time"`
	modified                string    `jsno:"modified" db:"modified"`
	UserID                  int16     `json:"user_id" db:"user_id"`
	WaiterID                int16     `jsno:"waiter_id" db:"waiter_id"`
	OutletID                int16     `json:"outlet_id" db:"outlet_id"`
	OrderStatus             int16     `jsno:"order_status" db:"order_status"`
	OrderType               int16     `jsno:"order_type" db:"order_type"`
	DelStatus               *string   `json:"del_status" db:"del_status"`
	IDFromOutlet            int16     `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload            bool      `json:"is_data_upload" db:"is_data_upload"`
	DateUpload              time.Time `json:"date_upload" db:"date_upload"`
}

// TblSalesDetails :
type TblSalesDetails struct {
	ID                       int16     `json:"id" db:"id"`
	FoodMenuID               int16     `json:"food_menu_id" db:"food_menu_id"`
	MenuName                 string    `json:"menu_name" db:"menu_name"`
	Qty                      int16     `json:"qty" db:"qty"`
	MenuPriceWithoutDiscount float32   `json:"menu_price_without_discount" db:"menu_price_without_discount"`
	MenuPriceWithDiscount    float32   `json:"menu_price_with_discount" db:"menu_price_with_discount"`
	MenuUnitPricec           float32   `json:"menu_unit_price" db:"menu_unit_price"`
	MenuVatPercentage        float32   `json:"menu_vat_percentage" db:"menu_vat_percentage"`
	MenuDiscountValue        string    `json:"menu_discount_value" db:"menu_discount_value"`
	DiscountType             string    `json:"discount_type" db:"discount_type"`
	MenuNote                 string    `json:"menu_note" db:"menu_note"`
	DiscountAmount           float64   `json:"discount_amount" db:""`
	ItemType                 string    `json:"item_type" db:"item_type"`
	CookingStatus            string    `json:"cooking_status" db:"cooking_status"`
	CookingStartTime         time.Time `json:"cooking_start_time" db:"cooking_start_time"`
	CookingDoneTime          time.Time `json:"cooking_done_time" db:"cooking_done_time"`
	PreviousID               int16     `json:"previous_id" db:"previous_id"`
	SalesID                  int16     `json:"sales_id" db:"sales_id"`
	OrderStatus              int16     `json:"order_status" db:"order_status"`
	UserID                   int16     `json:"user_id" db:"user_id"`
	OutletID                 int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus                *string   `json:"del_status" db:"del_status"`
	IDFromOutlet             int16     `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload             bool      `json:"is_data_upload" db:"is_data_upload"`
	DateUpload               time.Time `json:"date_upload" db:"date_upload"`
}

// TblSalesDetailsModifiers :
type TblSalesDetailsModifiers struct {
	ID             int16      `json:"id" db:"id"`
	ModifierID     int16      `json:"modifier_id" db:"modifier_id"`
	ModifierPrice  float32    `json:"modifier_price" db:"modifier_price"`
	FoodMenuID     int16      `json:"food_menu_id" db:"food_menu_id"`
	SalesID        int16      `json:"sales_id" db:"sales_id"`
	OrderStatus    int16      `json:"order_status" db:"order_status"`
	SalesDetailsID int16      `json:"sales_details_id" db:"sales_details_id"`
	UserID         int16      `json:"user_id" db:"user_id"`
	OutletID       int16      `json:"outlet_id" db:"outlet_id"`
	CustomerID     int16      `json:"customer_id" db:"customer_id"`
	IDFromOutlet   *int16     `json:"id_from_outlet" db:"id_from_outlet"`
	IsDataUpload   *int16     `json:"is_data_upload" db:"is_data_upload"`
	DateUpload     *time.Time `json:"date_upload" db:"date_upload"`
}

// TblSaleConsumptions :
type TblSaleConsumptions struct {
	ID           int16      `json:"id" db:"id"`
	SaleID       *int16     `json:"sale_id" db:"sale_id"`
	OrderStatus  int16      `json:"order_status" db:"order_status"`
	UserID       *int16     `json:"user_id" db:"user_id"`
	OutletID     *int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus    *string    `json:"del_status" db:"del_status"`
	IDFromOutlet *int16     `json:"id_from_outlet" db:"id_from_outlet"`
	IsDataUpload *int16     `json:"is_data_upload" db:"is_data_upload"`
	DateUpload   *time.Time `json:"date_upload" db:"date_upload"`
}

// TblSaleDataConsumptions :
type TblSaleDataConsumptions struct {
	ID                                    int16                                    `json:"id" db:"id"`
	SaleID                                *int16                                   `json:"sale_id" db:"sale_id"`
	OrderStatus                           int16                                    `json:"order_status" db:"order_status"`
	UserID                                *int16                                   `json:"user_id" db:"user_id"`
	OutletID                              *int16                                   `json:"outlet_id" db:"outlet_id"`
	DelStatus                             *string                                  `json:"del_status" db:"del_status"`
	IDFromOutlet                          *int16                                   `json:"id_from_outlet" db:"id_from_outlet"`
	IsDataUpload                          *int16                                   `json:"is_data_upload" db:"is_data_upload"`
	DateUpload                            *time.Time                               `json:"date_upload" db:"date_upload"`
	TblSaleConsumptionsOfMenus            []*TblSaleConsumptionsOfMenus            `json:"TblSaleConsumptionsOfMenus"`
	TblSaleConsumptionsOfModifiersOfMenus []*TblSaleConsumptionsOfModifiersOfMenus `json:"TblSaleConsumptionsOfModifiersOfMenus"`
}

// TblSaleConsumptionsOfMenus :
type TblSaleConsumptionsOfMenus struct {
	ID                int16      `json:"id" db:"id"`
	IngredientID      *int16     `json:"ingredient_id" db:"ingredient_id"`
	Consumption       *float32   `json:"consumption" db:"consumption"`
	SaleConsumptionID *int16     `json:"sale_consumption_id" db:"sale_consumption_id"`
	SalesID           int16      `json:"sales_id" db:"sales_id"`
	OrderStatus       int16      `json:"order_status" db:"order_status"`
	FoodMenuID        int16      `json:"food_menu_id" db:"food_menu_id"`
	UserID            *int16     `json:"user_id" db:"user_id"`
	OutletID          *int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus         *string    `json:"del_status" db:"del_status"`
	IDFromOutlet      *int16     `json:"id_from_outlet" db:"id_from_outlet"`
	IsDataUpload      *int16     `json:"is_data_upload" db:"is_data_upload"`
	DateUpload        *time.Time `json:"date_upload" db:"date_upload"`
}

// TblSaleConsumptionsOfModifiersOfMenus :
type TblSaleConsumptionsOfModifiersOfMenus struct {
	ID                int16      `json:"id" db:"id"`
	IngredientID      *int16     `json:"ingredient_id" db:"ingredient_id"`
	Consumption       *float32   `json:"consumption" db:"consumption"`
	SaleConsumptionID *int16     `json:"sale_consumption_id" db:"sale_consumption_id"`
	SalesID           int16      `json:"sales_id" db:"sales_id"`
	OrderStatus       int16      `json:"order_status" db:"order_status"`
	FoodMenuID        int16      `json:"food_menu_id" db:"food_menu_id"`
	UserID            *int16     `json:"user_id" db:"user_id"`
	OutletID          *int16     `json:"outlet_id" db:"outlet_id"`
	DelStatus         *string    `json:"del_status" db:"del_status"`
	IDFromOutlet      *int16     `json:"id_from_outlet" db:"id_from_outlet"`
	IsDataUpload      *int16     `json:"is_data_upload" db:"is_data_upload"`
	DateUpload        *time.Time `json:"date_upload" db:"date_upload"`
}

// TblDataSalesDetails :
type TblDataSalesDetails struct {
	ID                       int16                       `json:"id" db:"id"`
	FoodMenuID               int16                       `json:"food_menu_id" db:"food_menu_id"`
	MenuName                 string                      `json:"menu_name" db:"menu_name"`
	Qty                      int16                       `json:"qty" db:"qty"`
	MenuPriceWithoutDiscount float32                     `json:"menu_price_without_discount" db:"menu_price_without_discount"`
	MenuPriceWithDiscount    float32                     `json:"menu_price_with_discount" db:"menu_price_with_discount"`
	MenuUnitPricec           float32                     `json:"menu_unit_price" db:"menu_unit_price"`
	MenuVatPercentage        float32                     `json:"menu_vat_percentage" db:"menu_vat_percentage"`
	MenuDiscountValue        string                      `json:"menu_discount_value" db:"menu_discount_value"`
	DiscountType             string                      `json:"discount_type" db:"discount_type"`
	MenuNote                 string                      `json:"menu_note" db:"menu_note"`
	DiscountAmount           float64                     `json:"discount_amount" db:""`
	ItemType                 string                      `json:"item_type" db:"item_type"`
	CookingStatus            string                      `json:"cooking_status" db:"cooking_status"`
	CookingStartTime         time.Time                   `json:"cooking_start_time" db:"cooking_start_time"`
	CookingDoneTime          time.Time                   `json:"cooking_done_time" db:"cooking_done_time"`
	PreviousID               int16                       `json:"previous_id" db:"previous_id"`
	SalesID                  int16                       `json:"sales_id" db:"sales_id"`
	OrderStatus              int16                       `json:"order_status" db:"order_status"`
	UserID                   int16                       `json:"user_id" db:"user_id"`
	OutletID                 int16                       `json:"outlet_id" db:"outlet_id"`
	DelStatus                *string                     `json:"del_status" db:"del_status"`
	IDFromOutlet             int16                       `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload             bool                        `json:"is_data_upload" db:"is_data_upload"`
	DateUpload               time.Time                   `json:"date_upload" db:"date_upload"`
	TblSalesDetailsModifiers []*TblSalesDetailsModifiers `json:"TblSalesDetailsModifiers"`
}

// TblDataSales :
type TblDataSales struct {
	ID                      int16                    `json:"id" db:"id"`
	CustomerID              string                   `jsno:"customer_id" db:"customer_id"`
	SaleNo                  string                   `jsno:"sale_no" db:"sale_no"`
	TotalItems              int16                    `jsno:"total_items" db:"total_items"`
	SubTotal                float32                  `jsno:"sub_total" db:"sub_total"`
	PaidAmount              float64                  `jsno:"paid_amount" db:"paid_amount"`
	DueAmount               float32                  `jsno:"due_amount" db:"due_amount"`
	Disc                    string                   `jsno:"disc" db:"disc"`
	DiscActual              float32                  `jsno:"disc_actual" db:"disc_actual"`
	Vat                     float32                  `jsno:"vat" db:"vat"`
	TotalPayable            float32                  `jsno:"total_payable" db:"total_payable"`
	PaymentMethodID         int16                    `jsno:"payment_method_id" db:"payment_method_id"`
	CloseTime               time.Time                `jsno:"close_time" db:"close_time"`
	TableID                 int16                    `jsno:"table_id" db:"table_id"`
	TotalItemDiscountAmount float32                  `jsno:"total_item_discount_amount" db:"total_item_discount_amount"`
	SubTotalWithDiscount    float32                  `jsno:"sub_total_with_discount" db:"sub_total_with_discount"`
	SubTotalDiscountAmount  float32                  `jsno:"sub_total_discount_amount" db:"sub_total_discount_amount"`
	TotalDiscountAmount     float32                  `jsno:"total_discount_amount" db:"total_discount_amount"`
	DeliveryCharge          float32                  `jsno:"delivery_charge" db:"delivery_charge"`
	SubTotalDiscountValue   string                   `jsno:"sub_total_discount_value" db:"sub_total_discount_value"`
	SubTotalDiscountType    string                   `jsno:"sub_total_discount_type" db:"sub_total_discount_type"`
	SaleDate                string                   `jsno:"sale_date" db:"sale_date"`
	DateTime                time.Time                `jsno:"date_time" db:"date_time"`
	OrderTime               time.Time                `jsno:"order_time" db:"order_time"`
	CookingStartTime        time.Time                `jsno:"cooking_start_time" db:"cooking_start_time"`
	CookingDoneTime         time.Time                `jsno:"cooking_done_time" db:"cooking_done_time"`
	modified                string                   `jsno:"modified" db:"modified"`
	UserID                  int16                    `json:"user_id" db:"user_id"`
	WaiterID                int16                    `jsno:"waiter_id" db:"waiter_id"`
	OutletID                int16                    `json:"outlet_id" db:"outlet_id"`
	OrderStatus             int16                    `jsno:"order_status" db:"order_status"`
	OrderType               int16                    `jsno:"order_type" db:"order_type"`
	DelStatus               *string                  `json:"del_status" db:"del_status"`
	IDFromOutlet            int16                    `json:"id_from_outlet" db:"id_from_outlet"`
	ISDataUpload            bool                     `json:"is_data_upload" db:"is_data_upload"`
	DateUpload              time.Time                `json:"date_upload" db:"date_upload"`
	TblDataSalesDetails     []*TblDataSalesDetails   `json:"TblDataSalesDetails"`
	TblSaleDataConsumptions *TblSaleDataConsumptions `json:"TblSaleConsumptions"`
}
