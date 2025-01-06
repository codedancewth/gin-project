package main

import (
	"gin-project/mysql"
	"gin-project/post"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	(&mysql.Mysql{UserName: mysql.UserName, PassWord: mysql.PassWord}).InitMysql()

	// 构建一个gin的引擎实例.
	r := gin.Default()

	// 定义 路由
	r.POST("/describe_user_list", post.DescribeUserInfo)
	r.POST("/create_user", post.CreateUserInfo)

	r.Run()
}
