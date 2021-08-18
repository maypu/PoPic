package database

import (
	"PoPic/model"
	"PoPic/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysql() *gorm.DB {
	user := utils.GetConfig("mysql.user")
	password := utils.GetConfig("mysql.password")
	domain := utils.GetConfig("mysql.domain")
	port := utils.GetConfig("mysql.port")
	dbname := utils.GetConfig("mysql.dbname")
	charset := utils.GetConfig("mysql.charset")
	prefix := utils.GetConfig("mysql.prefix")

	dsn := user + ":" + password + "@tcp(" + domain + ":" + port + ")/" + dbname + "?charset=" + charset + "&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix, // 表名前缀
		},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("failed to connect database")
		db.AutoMigrate(&model.User{}, &model.Platform{}, &model.Upload{})
		//return nil
	}
	//defer db.Close()
	return db
}
