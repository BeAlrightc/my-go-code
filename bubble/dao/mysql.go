package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化一个全局的变量DB
var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(192.168.20.10:3306)/bubble?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功!")
	return DB.DB().Ping()
}
func Close() {
	DB.Close()
}
