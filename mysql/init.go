package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	UserName = "root"   // 用户的账号
	PassWord = "123456" // 用户的密码
)

type Mysql struct {
	UserName string
	PassWord string
}

var MysqlConnect *gorm.DB

// InitMysql 使用 GORM 初始化连接到 MySQL 数据库并执行查询
func (m *Mysql) InitMysql() *gorm.DB {
	fmt.Println("start")
	// 数据库连接信息
	// 链接格式【username:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local】
	// username 用户的账号
	// password 用户的密码
	// dbname 指向的库
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/project?charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.PassWord)

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	s, err := db.DB()

	if err != nil {
		log.Fatalf("set sql conn err [%v]", err)
	}
	s.SetMaxOpenConns(10) // 设置最大打开连接数
	s.SetMaxIdleConns(5)  // 设置最大空闲连接数

	// 注意使用db.bug的方法，会导致日志的数据，影响系统的性能
	//db = db.Debug()

	MysqlConnect = db
	return db
}
