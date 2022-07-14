package model

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

func Database(connsting string) {
	db, err := gorm.Open("mysql", connsting)
	if err != nil {
		panic(err)
	}
	db.LogMode(true) //开启日志
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)       //表明不加s
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
