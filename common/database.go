package common

import (
	"database/sql"
	"fmt"
	"time"

	//注册驱动器 _下划线表示执行驱动中的init函数，不使用其他函数
	_ "github.com/go-sql-driver/mysql"
)

// 创建数据库连接
func ConnMySQL() *sql.DB {
	// 数据源名
	driverName := "mysql"
	root := "root"
	passworld := "123456"
	local := "localhost"
	prod := "3306"
	dbName := "chat-web"
	// 用户名root
	// 密码1234
	// tcp协议连接
	// 数据库地址
	// 数据库 wms
	dataSourceName := root + ":" + passworld + "@" + "tcp" + "(" + local + ":" + prod + ")" + "/" + dbName
	fmt.Println(dataSourceName, "dataSourceName")
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// 连接测试
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	fmt.Println("数据库连接成功")
	return db
}
