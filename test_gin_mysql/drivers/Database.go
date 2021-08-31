package drivers

import (
	"fmt"
	"github.com/yourbrainrun/test_go/test_gin_mysql/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var databaseInstance *gorm.DB
var databaseOnce sync.Once

func GetDatabase() *gorm.DB {
	databaseOnce.Do(func() {
		conf := configs.GetDatabaseConfig()
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
		)

		var err error
		databaseInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}

		sqlDb, err := databaseInstance.DB()
		if err != nil {
			panic(err)
		}
		sqlDb.SetMaxIdleConns(10)
		sqlDb.SetMaxOpenConns(100)
	})
	return databaseInstance
}
