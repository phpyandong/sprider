package config

import (
	"log"
	"os"
	"sync"
	"time"
	"github.com/jinzhu/gorm"
)

const (
	ProjectName   = "Sprider"
	SecureLineKey = "#sprider#"
)

var ReadDB *gorm.DB
var WriteDB *gorm.DB
var once sync.Once

/*
  初始化数据库
*/
func InitDB() {
	once.Do(func() {
		initReadDB()
		initWriteDB()
	})
}


/*
 初始化读库
*/
func initReadDB() {
	env := os.Getenv("ENV")
	isLog := true
	if env == "pro" {
		isLog = false
	}
	ReadDB, _ = getDatabaseConnection(*AppSetting.DatabaseConfig.DatabaseOptions[0].ReadDBConns[0], isLog)
}


/*
 初始化写库
*/
func initWriteDB() {
	isLog := true
	env := os.Getenv("ENV")
	if env == "pro" {
		isLog = false
	}
	WriteDB, _ = getDatabaseConnection(*AppSetting.DatabaseConfig.DatabaseOptions[0].WirteDBConns[0], isLog)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 数据库链接
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func getDatabaseConnection(connectionOption DatabaseConnectionOption, isLog bool) (*gorm.DB, error) {
	dsn := connectionOption.Username + ":" + connectionOption.Password + "@tcp(" + connectionOption.Host + ")/" + connectionOption.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	dbMap, err := gorm.Open(connectionOption.Dialect, dsn)

	if err != nil {
		log.Printf("Error connecting to db: %s", err.Error())
	}
	dbMap.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
	dbMap.DB().SetMaxIdleConns(16)
	dbMap.DB().SetMaxOpenConns(80)
	dbMap.DB().SetConnMaxLifetime(time.Hour)
	dbMap.DB().Ping()
	dbMap.LogMode(isLog)

	return dbMap, err
}
