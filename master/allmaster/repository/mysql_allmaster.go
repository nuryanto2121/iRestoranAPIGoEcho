package repoallmaster

import (
	"context"

	"github.com/jmoiron/sqlx"

	allmaster "yanto/irestora/master/allmaster/domain"

	"yanto/irestora/models"
)

//mysqlAllMasterRepository :
type mysqlAllMasterRepository struct {
	Conn *sqlx.DB
}

// NewMysqlAllMasterRepository :
func NewMysqlAllMasterRepository(Conn *sqlx.DB) allmaster.Repository {
	return &mysqlAllMasterRepository{Conn}
}

func (m *mysqlAllMasterRepository) GetAllMaster(ctx context.Context) (result models.AllMaster, err error) {
	var (
		IngredientCategories []models.IngredientCategories
		Units                []models.Units
		Ingredients          []models.Ingredients
		Vats                 []models.Vats
		Modifiers            []models.Modifiers
		ModifiersIngredients []models.ModifiersIngredients
		FoodMenuCategories   []models.FoodMenuCategories
		FoodMenus            []models.FoodMenus
		FoodMenusIngredients []models.FoodMenusIngredients
		FoodMenusModifiers   []models.FoodMenusModifiers
		Suppliers            []models.Suppliers
		Customers            []models.Customers
		ExpenseItems         []models.ExpenseItems
		Employees            []models.Employees
		PaymentMethod        []models.PaymentMethod
		Tables               []models.Tables
		RuleDiscount         []models.RuleDiscount
	)

	// getList IngredientCategories
	err = m.Conn.SelectContext(ctx, &IngredientCategories, allmaster.QueryIngredientsCategories)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Units
	err = m.Conn.SelectContext(ctx, &Units, allmaster.QueryUnits)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Ingredients
	err = m.Conn.SelectContext(ctx, &Ingredients, allmaster.QueryIngredients)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Vats
	err = m.Conn.SelectContext(ctx, &Vats, allmaster.QueryVats)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Modifiers
	err = m.Conn.SelectContext(ctx, &Modifiers, allmaster.QueryModifiers)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList ModifiersIngredients
	err = m.Conn.SelectContext(ctx, &ModifiersIngredients, allmaster.QueryModifiersIngredients)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList FoodMenuCategories
	err = m.Conn.SelectContext(ctx, &FoodMenuCategories, allmaster.QueryFoodMenuCategories)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList FoodMenus
	err = m.Conn.SelectContext(ctx, &FoodMenus, allmaster.QueryFoodMenus)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList FoodMenusIngredients
	err = m.Conn.SelectContext(ctx, &FoodMenusIngredients, allmaster.QueryFoodMenusIngredients)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList FoodMenusModifiers
	err = m.Conn.SelectContext(ctx, &FoodMenusModifiers, allmaster.QueryFoodMenusModifiers)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Suppliers
	err = m.Conn.SelectContext(ctx, &Suppliers, allmaster.QuerySuppliers)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Customers
	err = m.Conn.SelectContext(ctx, &Customers, allmaster.QueryCustomers)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList ExpenseItems
	err = m.Conn.SelectContext(ctx, &ExpenseItems, allmaster.QueryExpenseItems)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Employees
	err = m.Conn.SelectContext(ctx, &Employees, allmaster.QueryEmployees)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList PaymentMethod
	err = m.Conn.SelectContext(ctx, &PaymentMethod, allmaster.QueryPaymentMethod)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList Tables
	err = m.Conn.SelectContext(ctx, &Tables, allmaster.QueryTables)
	if err != nil {
		return models.AllMaster{}, err
	}

	// getList RuleDiscount
	err = m.Conn.SelectContext(ctx, &RuleDiscount, allmaster.QueryRuleDiscount)
	if err != nil {
		return models.AllMaster{}, err
	}
	result.IngredientCategories = IngredientCategories
	result.Units = Units
	result.Ingredients = Ingredients
	result.Vats = Vats
	result.Modifiers = Modifiers
	result.ModifiersIngredients = ModifiersIngredients
	result.FoodMenuCategories = FoodMenuCategories
	result.FoodMenus = FoodMenus
	result.FoodMenusIngredients = FoodMenusIngredients
	result.FoodMenusModifiers = FoodMenusModifiers
	result.Suppliers = Suppliers
	result.Customers = Customers
	result.ExpenseItems = ExpenseItems
	result.Employees = Employees
	result.PaymentMethod = PaymentMethod
	result.Tables = Tables
	result.RuleDiscount = RuleDiscount

	return result, nil
}
