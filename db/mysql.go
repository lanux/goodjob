package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lanux/goodjob/v1/config"
	"time"
)

var (
	mysql *gorm.DB
)

func init() {
	c := config.Mysql
	connArgs := fmt.Sprintf("%s:%s@%s", c.Username, c.Password, c.Url)
	MYSQL, err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic(err)
	}
	MYSQL.DB().SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))
	MYSQL.DB().SetMaxIdleConns(c.MaxIdle)
	MYSQL.DB().SetMaxOpenConns(c.MaxOpenConns)
	mysql = MYSQL
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Instance() *gorm.DB {
	return mysql
}
