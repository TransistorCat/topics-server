package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initMySQL() error {
	dataSourceName := "root:root@tcp(127.0.0.1:3306)/TopicServerDB?charset=utf8&parseTime=True"
	var err error
	if DB, err = gorm.Open(mysql.Open(dataSourceName), nil); err != nil {
		return err
	}
	return nil
}
