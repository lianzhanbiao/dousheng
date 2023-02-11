package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	log.Println("开始初始化数据库连接...")
	//todo: 使用配置config连接
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "root:123456@tcp(150.158.79.245:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("连接失败")
	}
	log.Println("开始初始化用户表...")
	db.AutoMigrate(&(User{}))
	return db
}
