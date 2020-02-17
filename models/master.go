package models

import (
	"time"
)

//IngredientCategories :
type IngredientCategories struct {
	ID           int16   `json:"id" db:"id"`
	CategoryName *string `json:"category_name" db:"category_name"`
	Description  *string `json:"description" db:"description"`
	UserID       int16   `json:"user_id" db:"user_id"`
	CompanyID    int16   `json:"company_id" db:"company_id"`
	DelStatus    *string `json:"del_status" db:"del_status"`
}

// Units :
type Units struct {
	ID          int16   `json:"id" db:"id"`
	UnitName    *string `json:"unit_name" db:"unit_name"`
	Description *string `json:"description" db:"description"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

// Ingredients :
type Ingredients struct {
	ID            int16   `json:"id" db:"id"`
	Code          *string `json:"code" db:"code"`
	Name          *string `json:"name" db:"name"`
	CategoryID    int16   `json:"category_id" db:"category_id"`
	PurchasePrice float32 `json:"purchase_price" db:"purchase_price"`
	AlertQuantity float32 `json:"alert_quantity" db:"alert_quantity"`
	UnitID        int16   `json:"unit_id" db:"unit_id"`
	UserID        int16   `json:"user_id" db:"user_id"`
	CompanyID     int16   `json:"company_id" db:"company_id"`
	DelStatus     *string `json:"del_status" db:"del_status"`
}

//Vats :
type Vats struct {
	ID         int16   `json:"id" db:"id"`
	Name       *string `json:"name" db:"name"`
	Percentage float32 `json:"percentage" db:"percentage"`
	UserID     int16   `json:"user_id" db:"user_id"`
	CompanyID  int16   `json:"company_id" db:"company_id"`
	DelStatus  *string `json:"del_status" db:"del_status"`
}

//Modifiers :
type Modifiers struct {
	ID          int16   `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Price       float32 `json:"price" db:"price"`
	Description *string `json:"description" db:"description"`
	UserID      int16   `json:"user_id" db:"user_id"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

//ModifiersIngredients :
type ModifiersIngredients struct {
	ID           int16   `json:"id" db:"id"`
	IngredientID int16   `json:"ingredient_id" db:"ingredient_id"`
	Consumption  float32 `json:"consumption" db:"consumption"`
	ModifierID   *string `json:"modifier_id" db:"modifier_id"`
	UserID       int16   `json:"user_id" db:"user_id"`
	CompanyID    int16   `json:"company_id" db:"company_id"`
	DelStatus    *string `json:"del_status" db:"del_status"`
}

//FoodMenuCategories :
type FoodMenuCategories struct {
	ID           int16   `json:"id" db:"id"`
	CategoryName *string `json:"category_name" db:"category_name"`
	Description  *string `json:"description" db:"description"`
	UserID       int16   `json:"user_id" db:"user_id"`
	CompanyID    int16   `json:"company_id" db:"company_id"`
	DelStatus    *string `json:"del_status" db:"del_status"`
}

//FoodMenus :
type FoodMenus struct {
	ID           int16   `json:"id" db:"id"`
	Code         *string `json:"code" db:"code"`
	Name         *string `json:"name" db:"name"`
	CategoryID   int16   `json:"category_id" db:"category_id"`
	Description  *string `json:"description" db:"description"`
	SalePrice    float32 `json:"sale_price" db:"sale_price"`
	VatID        int16   `json:"vat_id" db:"vat_id"`
	UserID       int16   `json:"user_id" db:"user_id"`
	CompanyID    int16   `json:"company_id" db:"company_id"`
	Photo        *string `json:"photo" db:"photo"`
	VegItem      *string `json:"veg_item" db:"veg_item"`
	BeverageItem *string `json:"beverage_item" db:"beverage_item"`
	BarItem      *string `json:"bar_item" db:"bar_item"`
	DelStatus    *string `json:"del_status" db:"del_status"`
}

//FoodMenusIngredients :
type FoodMenusIngredients struct {
	ID           int16   `json:"id" db:"id"`
	IngredientID int16   `json:"ingredient_id" db:"ingredient_id"`
	Consumption  float32 `json:"consumption" db:"consumption"`
	FoodMenuID   int32   `json:"food_menu_id" db:"food_menu_id"`
	UserID       int16   `json:"user_id" db:"user_id"`
	CompanyID    int16   `json:"company_id" db:"company_id"`
	DelStatus    *string `json:"del_status" db:"del_status"`
}

//FoodMenusModifiers :
type FoodMenusModifiers struct {
	ID         int16   `json:"id" db:"id"`
	ModifierID int16   `json:"modifier_id" db:"modifier_id"`
	FoodMenuID int16   `json:"food_menu_id" db:"food_menu_id"`
	UserID     int16   `json:"user_id" db:"user_id"`
	OutletID   int16   `json:"outlet_id" db:"outlet_id"`
	CompanyID  int16   `json:"company_id" db:"company_id"`
	DelStatus  *string `json:"del_status" db:"del_status"`
}

//Suppliers :
type Suppliers struct {
	ID            int16   `json:"id" db:"id"`
	Name          *string `json:"name" db:"name"`
	ContactPerson *string `json:"contact_person" db:"contact_person"`
	Phone         *string `json:"phone" db:"phone"`
	Email         *string `json:"email" db:"email"`
	Address       *string `json:"address" db:"address"`
	Description   *string `json:"description" db:"description"`
	UserID        int16   `json:"user_id" db:"user_id"`
	CompanyID     int16   `json:"company_id" db:"company_id"`
	DelStatus     *string `json:"del_status" db:"del_status"`
}

//Customers :
type Customers struct {
	ID                int16      `json:"id" db:"id"`
	Name              *string    `json:"name" db:"name"`
	Phone             *string    `json:"phone" db:"phone"`
	Email             *string    `json:"email" db:"email"`
	Address           *string    `json:"address" db:"address"`
	AreaID            int16      `json:"area_id" db:"area_id"`
	UserID            int16      `json:"user_id" db:"user_id"`
	CompanyID         int16      `json:"company_id" db:"company_id"`
	DelStatus         *string    `json:"del_status" db:"del_status"`
	DateOfBirth       *time.Time `json:"date_of_birth" db:"date_of_birth"`
	DateOfAnniversary *time.Time `json:"date_of_anniversary" db:"date_of_anniversary"`
}

//ExpenseItems :
type ExpenseItems struct {
	ID          int16   `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	UserID      int16   `json:"user_id" db:"user_id"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

//Employees :
type Employees struct {
	ID          int16   `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Designation *string `json:"designation" db:"designation"`
	Phone       *string `json:"phone" db:"phone"`
	Description *string `json:"description" db:"description"`
	UserID      int16   `json:"user_id" db:"user_id"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

//PaymentMethod :
type PaymentMethod struct {
	ID          int16   `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	UserID      int16   `json:"user_id" db:"user_id"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

//Tables :
type Tables struct {
	ID          int16   `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	SitCapacity *string `json:"sit_capacity" db:"sit_capacity"`
	Position    *string `json:"position" db:"position"`
	Description *string `json:"description" db:"description"`
	UserID      int16   `json:"user_id" db:"user_id"`
	OutletID    int16   `json:"outlet_id" db:"outlet_id"`
	CompanyID   int16   `json:"company_id" db:"company_id"`
	DelStatus   *string `json:"del_status" db:"del_status"`
}

// RuleDiscount :
type RuleDiscount struct {
	ID              int16     `json:"id" db:"id"`
	CompanyID       int16     `json:"company_id" db:"company_id"`
	OutletID        int16     `json:"outlet_id" db:"outlet_id"`
	Customer        *string   `json:"customer" db:"customer"`
	StartDate       time.Time `json:"start_date" db:"start_date"`
	EndDate         time.Time `json:"end_date" db:"end_date"`
	MinAmount       float32   `json:"min_amount" db:"min_amount"`
	DiscountType    *string   `json:"discount_type" db:"discount_type"`
	Amount          float32   `json:"amount" db:"amount"`
	Percent         int16     `json:"percent" db:"percent"`
	CurrencyCd      *string   `json:"currency_cd" db:"currency_cd"`
	TotalCategory   *int16    `json:"total_category" db:"total_category"`
	MenuType        *string   `json:"menu_type" db:"menu_type"`
	MenuCategoryIDs *string   `json:"menu_category_ids" db:"menu_category_ids"`
	MenuIDs         *string   `json:"menu_ids" db:"menu_ids"`
	UserInput       *string   `json:"user_input" db:"user_input"`
	DateInput       time.Time `json:"date_input" db:"date_input"`
	UserEdit        *string   `json:"user_edit" db:"user_edit"`
	DateEdit        time.Time `json:"date_edit" db:"date_edit"`
	DelStatus       *string   `json:"del_status" db:"del_status"`
}
