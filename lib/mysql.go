package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var db *gorm.DB

func MysqlConnect() *gorm.DB{
	if nil!=db{
		return db
	}
	dsn := "root:mYsqlPwd!@tcp(139.9.35.173:3360)/shopx-core?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err!=nil{
		panic(err)
	}
	return db
}