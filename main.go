package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"yanto/irestora/middleware"
	dbgorm "yanto/irestora/pkg/db_gorm"
	"yanto/irestora/pkg/setting"

	_ingeredientcategoriesDelivery "yanto/irestora/master/ingeredient_categories/delivery"
	_ingeredientcategoriesRepo "yanto/irestora/master/ingeredient_categories/repository"
	_ingeredientcategoriesUsecase "yanto/irestora/master/ingeredient_categories/usecase"

	_allmasterDelivery "yanto/irestora/master/allmaster/delivery"
	_allmasterRepo "yanto/irestora/master/allmaster/repository"
	_allmasterUsecase "yanto/irestora/master/allmaster/usecase"

	_purchaseDelivery "yanto/irestora/transaksi/purchase/delivery"
	_purchaseRepo "yanto/irestora/transaksi/purchase/repository"
	_purchaseUsecase "yanto/irestora/transaksi/purchase/usecase"
)

var db *gorm.DB

func init() {
	setting.Setup()
	db = dbgorm.InitGormDB()
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	dbConn2, err2 := sqlx.Open("mysql", dsn)
	if err != nil && viper.GetBool("debug") {
		fmt.Println(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err2 = dbConn2.Ping()
	if err != nil {
		log.Fatal(err2)
		os.Exit(1)
	}

	defer func() {
		err2 := dbConn2.Close()
		if err != nil {
			log.Fatal(err2)
		}
	}()

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	// authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	// ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	/*start ingeredient*/
	icRepo := _ingeredientcategoriesRepo.NewMysqlIngeredientCategoriesRepository(dbConn2)
	icUsecase := _ingeredientcategoriesUsecase.NewIngeredientCategoryUsecase(icRepo, timeoutContext)
	_ingeredientcategoriesDelivery.NewIngeredientCategoriesHandler(e, icUsecase)
	/*end ingeredient*/
	// au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	// _articleHttpDeliver.NewArticleHandler(e, au)

	/*start all master*/
	allMasterRepo := _allmasterRepo.NewMysqlAllMasterRepository(dbConn2)
	allMasterUsecase := _allmasterUsecase.NewallMasterUsecase(allMasterRepo, timeoutContext)
	_allmasterDelivery.NewAllMasterHandler(e, allMasterUsecase)
	/*end all master*/

	/*start purchase*/
	purchaseRepo := _purchaseRepo.NewMysqlPurchaseRepository(db)
	purchaseUsecase := _purchaseUsecase.NewPurchaseUsecase(purchaseRepo, timeoutContext)
	_purchaseDelivery.NewPurchaseHandler(e, purchaseUsecase)
	/*end purchase*/

	log.Fatal(e.Start(viper.GetString("server.address")))
}
