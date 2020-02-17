package setting

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

// Database :
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Name        string
	TablePrefix string
}

// DatabaseSetting :
var DatabaseSetting = &Database{}

// Server :
type Server struct {
	RunMode  string
	HTTPPort int
	TimeOut  int
}

//ServerSetting :
var ServerSetting = &Server{}

// Setup Load config from ini config.json
func Setup() {
	now := time.Now()
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

	DatabaseSetting.Type = viper.GetString(`database.type`)
	DatabaseSetting.Host = viper.GetString(`database.host`)
	DatabaseSetting.Port = viper.GetString(`database.port`)
	DatabaseSetting.User = viper.GetString(`database.user`)
	DatabaseSetting.Password = viper.GetString(`database.pass`)
	DatabaseSetting.Name = viper.GetString(`database.name`)
	DatabaseSetting.TablePrefix = viper.GetString(`database.table_prefix`)

	ServerSetting.RunMode = viper.GetString(`server.runmode`)
	ServerSetting.TimeOut = viper.GetInt(`server.timeout`)
	ServerSetting.HTTPPort = viper.GetInt(`server.http_port`)

	timeSpent := time.Since(now)
	log.Printf("Config setting is ready in %v", timeSpent)
}
