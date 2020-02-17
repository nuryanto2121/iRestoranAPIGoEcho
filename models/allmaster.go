package models

//AllMaster :
type AllMaster struct {
	IngredientCategories []IngredientCategories `json:"ingredient_categories"`
	Units                []Units                `json:"units"`
	Ingredients          []Ingredients          `json:"ingredients"`
	Vats                 []Vats                 `json:"vats"`
	Modifiers            []Modifiers            `json:"modifiers"`
	ModifiersIngredients []ModifiersIngredients `json:"modifiers_ingredients"`
	FoodMenuCategories   []FoodMenuCategories   `json:"food_menu_categories"`
	FoodMenus            []FoodMenus            `json:"food_menus"`
	FoodMenusIngredients []FoodMenusIngredients `json:"food_menus_ingredients"`
	FoodMenusModifiers   []FoodMenusModifiers   `json:"food_menus_modifiers"`
	Suppliers            []Suppliers            `json:"suppliers"`
	Customers            []Customers            `json:"customers"`
	ExpenseItems         []ExpenseItems         `json:"expense_items"`
	Employees            []Employees            `json:"employees"`
	PaymentMethod        []PaymentMethod        `json:"payment_method"`
	Tables               []Tables               `json:"tables"`
	RuleDiscount         []RuleDiscount         `json:"rule_discount"`
}
