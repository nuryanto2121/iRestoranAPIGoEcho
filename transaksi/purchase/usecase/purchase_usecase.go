package purchaseusecase

import (
	"context"
	"fmt"
	"time"
	"yanto/irestora/models"
	purchase "yanto/irestora/transaksi/purchase/domain"
	util "yanto/irestora/utils"
)

type purchaseUsecase struct {
	purchaseRepo   purchase.Repository
	contextTimeOut time.Duration
}

// NewPurchaseUsecase :
func NewPurchaseUsecase(a purchase.Repository, timeout time.Duration) purchase.Usecase {
	return &purchaseUsecase{
		purchaseRepo:   a,
		contextTimeOut: timeout,
	}
}

func (a *purchaseUsecase) GetPurchase(ctx context.Context) (result []*models.TblPurchase, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	listPurchase, err := a.purchaseRepo.GetPurchase()
	if err != nil {
		return nil, err
	}

	return listPurchase, nil
}

func (a *purchaseUsecase) AddPurchase(ctx context.Context, dp *models.DataPurchase) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	/*start header*/
	// x := struct{Foo string; Bar int }{"foo", 2}
	for _, dataPurchase := range dp.Purchase {

		PurchaseHeader := models.TblPurchase{}
		PurchaseHeader.ID = 0
		PurchaseHeader.ReferenceNo = dataPurchase.ReferenceNo
		PurchaseHeader.SupplierID = dataPurchase.SupplierID
		PurchaseHeader.Date = dataPurchase.Date
		PurchaseHeader.Subtotal = dataPurchase.Subtotal
		PurchaseHeader.Other = dataPurchase.Other
		PurchaseHeader.GrandTotal = dataPurchase.GrandTotal
		PurchaseHeader.Paid = dataPurchase.Paid
		PurchaseHeader.Due = dataPurchase.Due
		PurchaseHeader.Note = dataPurchase.Note
		PurchaseHeader.UserID = dataPurchase.UserID
		PurchaseHeader.OutletID = dataPurchase.OutletID
		PurchaseHeader.DelStatus = dataPurchase.DelStatus
		PurchaseHeader.IDFromOutlet = dataPurchase.ID
		PurchaseHeader.ISDataUpload = true
		PurchaseHeader.DateUpload = util.GetTimeNow()

		err = a.purchaseRepo.AddPurchase(&PurchaseHeader)
		if err != nil {
			return err
		}

		for _, dataPurchaseIngredient := range dataPurchase.TblPurchaseIngredients {
			PurchaseDetail := models.TblPurchaseIngredients{}
			PurchaseDetail.ID = 0
			PurchaseDetail.IngredientID = dataPurchaseIngredient.IngredientID
			PurchaseDetail.UnitPrice = dataPurchaseIngredient.UnitPrice
			PurchaseDetail.QuantityAmount = dataPurchaseIngredient.UnitPrice
			PurchaseDetail.Total = dataPurchaseIngredient.Total
			PurchaseDetail.PurchaseID = PurchaseHeader.ID
			PurchaseDetail.OutletID = dataPurchaseIngredient.OutletID
			PurchaseDetail.DelStatus = dataPurchaseIngredient.DelStatus
			PurchaseDetail.IDFromOutlet = dataPurchaseIngredient.ID
			PurchaseDetail.ISDataUpload = true
			PurchaseDetail.DateUpload = util.GetTimeNow()
			err = a.purchaseRepo.AddPurchaseIngeredient(&PurchaseDetail)
			if err != nil {
				return err
			}

		}

	}
	// v := reflect.ValueOf(dp).Elem()

	// values := make([]interface{}, v.NumField())

	// for i := 0; i < v.NumField(); i++ {
	// 	values[i] = v.Field(i).Interface()
	// }

	// fmt.Println(values)

	// for k, v := range values {
	// 	fmt.Println(k, "=>", v)
	// }
	// var datapostPurchase []*models.TbldataPurchase
	// datapostPurchase = dp.Purchase
	// for index, element := range dp {
	// 	fmt.Printf("%s", index)
	// 	fmt.Printf("%+v\n", element)
	// 	s := reflect.ValueOf(element)
	// 	for _, k := range s.MapKeys() {
	// 		fmt.Println(s.MapIndex(k))
	// 	}
	// fmt.Printf("%d", )
	// fmt.Println("Index :", index, " Element %v:", element)
	// fmt.Println("Element :", element[index])
	// }
	// m, ok := dp.(map[string]interface{})
	// if !ok {
	// 	return fmt.Errorf("want type map[string]interface{};  got %T", t)
	// }
	// jsonMap := make(map[string]interface{})
	// var dataPurchase models.DataPurchase
	// // var dataDetailpurchase []models.TblPurchaseIngredients
	// values := make([]interface{}, len(dp))
	// object := map[string]interface{}{}

	// for head, val := range dp {

	// 	fmt.Println(head, "=>", val)
	// 	// if head == "purchase" {

	// 	fmt.Printf("got %T", val)
	// 	b, err := json.Marshal(val)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	jsonstring := fmt.Sprintf("%s", string(b))
	// 	fmt.Print(jsonstring)
	// 	var a interface{}
	// 	err = json.Unmarshal(b, &a)
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 	}
	// 	fmt.Println(a)
	// 	fmt.Println(a.(*string))
	// 	s := reflect.ValueOf(a)
	// 	for _, k := range s.MapKeys() {
	// 		fmt.Println(s.MapIndex(k))
	// 	}

	// // Creating the maps for JSON
	// m := map[string]interface{}{}

	// // Parsing/Unmarshalling JSON encoding/json
	// err = json.Unmarshal([]byte(jsonstring), &m)
	// if err != nil {
	// 	panic(err)
	// }
	// parseMap(m)
	// b := []byte(val.(string))
	// Declared an empty interface
	// var result map[string]interface{}
	// json.Unmarshal(b, &result)
	// m, ok := val.(map[string]interface{})
	// if !ok {
	// 	return fmt.Errorf("want type map[string]interface{};  got %T", val)
	// }
	// fmt.Println(result["ingredient_id"])

	// for k, v := range m {
	// 	fmt.Println(k, "=>", v)
	// }
	// for h, v := range val {

	// }
	// fmt.Println(val[""])

	// }
	// err := json.Unmarshal()
	// if err != nil {
	// 	return e.JSON(http.StatusBadRequest, err.Error())
	// }
	// m, _ := v.(map[string]interface{})
	// for x, z := range v {
	// 	fmt.Println(x, "=>", z)
	// }

	// }

	/*end header*/

	return nil
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
