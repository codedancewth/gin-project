package post

import (
	"gin-project/mysql"
	"gin-project/mysql/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// DescribeUserInfo 获取用户基本信息
func DescribeUserInfo(c *gin.Context) {
	var user User
	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 查询用户的列表
	err, users := dao.GetUserList(mysql.MysqlConnect)
	if err != nil {
		log.Panicf("err [%v]", err)
		return
	}

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "User successful response",
		"user":    users,
	})
}
