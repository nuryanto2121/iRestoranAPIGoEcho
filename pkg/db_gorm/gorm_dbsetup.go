package dbgorm

import (
	"fmt"
	"log"
	"time"
	"yanto/irestora/pkg/setting"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitGormDB :
func InitGormDB() *gorm.DB {
	now := time.Now()
	var err error
	connectionParams := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name)
	// connectionParams := fmt.Sprintf(
	// 	"user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Name,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Port,
	// )

	db, err = gorm.Open(setting.DatabaseSetting.Type, connectionParams)
	if err != nil {
		// logging.Fatal("0", "models.Setup err: ", err)
		log.Printf("dbgorm.InitGormDB err : %v", err)
	}
	// defer db.Close()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	timeSpent := time.Since(now)
	log.Printf("Config database is ready in %v", timeSpent)
	return db
}
