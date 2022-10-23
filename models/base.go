package models

import (
	"github.com/colincheng/servicedesk/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/os/gfile"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
)

// 在其它model的实体类中可直接调用
var db *gorm.DB

func init() {
	db = InitDB()
}

func InitDB() *gorm.DB {
	var err error

	// sqlite3配置信息
	sqliteName := "servicedesk"

	dataSource := "database" + string(os.PathSeparator) + sqliteName
	if !gfile.Exists(dataSource) {
		if err := os.MkdirAll(path.Dir(dataSource), os.ModePerm); err != nil {
			helper.LogError(err.Error())
			panic("系统错误")
		}
		if _, err := os.Create(dataSource); err != nil {
			helper.LogError(err.Error())
			panic("系统错误")
		}
	}
	db, err = gorm.Open("sqlite3", dataSource)

	if err != nil {
		if err := db.Close(); err != nil {
			helper.LogError(err.Error())
			panic("系统错误")
		}
	}

	// 设置连接池，空闲连接
	db.DB().SetMaxIdleConns(50)
	// 打开链接
	db.DB().SetMaxOpenConns(100)

	// 严格匹配表名
	db.SingularTable(true)
	//打开debug，输出sql方便开发
	db.LogMode(true)

	return db
}
