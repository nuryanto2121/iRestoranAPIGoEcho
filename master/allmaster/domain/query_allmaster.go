package allmaster

const (
	//QueryIngredientsCategories :
	QueryIngredientsCategories = `SELECT id,category_name,description,user_id,company_id,del_status 
								FROM tbl_ingredient_categories WHERE del_status ='LIVE'`

	//QueryUnits :
	QueryUnits = `SELECT id,unit_name,description,company_id,del_status 
					FROM tbl_units WHERE del_status ='LIVE'`

	//QueryIngredients :
	QueryIngredients = `SELECT id,code,name,category_id,purchase_price,alert_quantity,unit_id,user_id,company_id,del_status 
						FROM tbl_ingredients  WHERE del_status ='LIVE'`

	// QueryVats :
	QueryVats = `SELECT id,name,percentage,user_id,company_id,del_status 
					FROM tbl_vats WHERE del_status ='LIVE'`

	// QueryModifiers :
	QueryModifiers = `SELECT id,name,price,description,user_id,company_id,del_status 
						FROM tbl_modifiers WHERE del_status ='LIVE'`

	//QueryModifiersIngredients :
	QueryModifiersIngredients = `SELECT id,ingredient_id,consumption,modifier_id,user_id,company_id,del_status 
								FROM tbl_modifier_ingredients WHERE del_status ='LIVE'`

	//QueryFoodMenuCategories :
	QueryFoodMenuCategories = `SELECT id,category_name,description,user_id,company_id,del_status 
							FROM tbl_food_menu_categories WHERE del_status ='LIVE'`

	// QueryFoodMenus :
	QueryFoodMenus = `SELECT id, code, name, category_id, description, sale_price, vat_id, user_id, company_id, photo, veg_item, beverage_item, bar_item, del_status 
						FROM tbl_food_menus WHERE del_status ='LIVE'`

	//QueryFoodMenusIngredients :
	QueryFoodMenusIngredients = `SELECT id, ingredient_id, consumption, food_menu_id, user_id, company_id, del_status 
									FROM tbl_food_menus_ingredients WHERE del_status ='LIVE'`

	//QueryFoodMenusModifiers :
	QueryFoodMenusModifiers = `SELECT
									id,
									modifier_id,
									food_menu_id,
									user_id,
									outlet_id,
									company_id,
									del_status
								FROM
									tbl_food_menus_modifiers
								WHERE del_status ='LIVE'`

	// QuerySuppliers :
	QuerySuppliers = `	SELECT
							id,
							name,
							contact_person,
							phone,
							email,
							address,
							description,
							user_id,
							company_id,
							del_status
						FROM
							tbl_suppliers
						WHERE del_status ='LIVE'`

	//QueryCustomers :
	QueryCustomers = `	SELECT
							id,
							name,
							phone,
							email,
							address,
							area_id,
							user_id,
							company_id,
							del_status,
							date_of_birth,
							date_of_anniversary
						FROM
							tbl_customers
						WHERE del_status ='LIVE'`

	//QueryExpenseItems :
	QueryExpenseItems = `SELECT
							id,
							name,
							description,
							user_id,
							company_id,
							del_status
						FROM
							tbl_expense_items
						WHERE del_status ='LIVE'`

	//QueryEmployees :
	QueryEmployees = `SELECT
						id,
						name,
						designation,
						phone,
						description,
						user_id,
						company_id,
						del_status
					FROM
						tbl_employees
					WHERE del_status ='LIVE'`

	//QueryPaymentMethod :
	QueryPaymentMethod = `	SELECT
								id,
								name,
								description,
								user_id,
								company_id,
								del_status
							FROM
								tbl_payment_methods
							WHERE del_status ='LIVE'`

	//QueryTables :
	QueryTables = `	SELECT
						id,
						name,
						sit_capacity,
						position,
						description,
						user_id,
						outlet_id,
						company_id,
						del_status
					FROM
						tbl_tables
					WHERE del_status ='LIVE'`

	// QueryRuleDiscount :
	QueryRuleDiscount = `	SELECT
								id,
								company_id,
								outlet_id,
								customer,
								start_date,
								end_date,
								min_amount,
								discount_type,
								amount,
								percent,
								currency_cd,
								total_category,
								menu_type,
								menu_category_ids,
								menu_ids,
								user_input,
								date_input,
								user_edit,
								date_edit,
								del_status
							FROM
								tbl_rule_discount
							WHERE del_status ='LIVE'`
)
