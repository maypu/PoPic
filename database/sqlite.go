package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Sqlite *gorm.DB

func InitSQLite() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "calcu_", // 表名前缀
		},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("failed to connect database")
		return nil
	}
	Sqlite = db
	//Sqlite.DB().SetMaxIdleConns(1000)
	//Sqlite.DB().SetMaxOpenConns(100000)
	//Sqlite.DB().SetConnMaxLifetime(-1)

	//defer Sqlite.Close()
	return Sqlite
}
