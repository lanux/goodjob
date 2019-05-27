package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goodjob/config"
	"time"
)

var (
	DB *gorm.DB
)

func init() {
	DB = new()
}

func new() *gorm.DB {
	c := config.Global.Mysql
	connArgs := fmt.Sprintf("%s:%s@%s", c.Username, c.Password, c.Url)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	db.DB().SetMaxIdleConns(c.MaxIdle)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	return db
}
