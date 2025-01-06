package post

import (
	"gin-project/mysql"
	"gin-project/mysql/dao"
	"gin-project/mysql/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 构建easy-post，用于测试gin框架的情况

type CreateUserParam struct {
	UserName     string `json:"user_name"`     // 用户姓名
	UserAccount  string `json:"user_account"`  // 用户账号
	UserPassword string `json:"user_password"` // 用户密码
	Status       int    `json:"status"`        // 是否活跃
	CreatedTime  int    `json:"created_time"`  // 创建时间
	UpdatedTime  int    `json:"updated_time"`  // 更新时间
	IsDeleted    int    `json:"is_deleted"`    // 是否已经删除
}

type CreateUserResp struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

// CreateUserInfo 创建用户基本信息
func CreateUserInfo(c *gin.Context) {
	var userInfo CreateUserParam
	// 解析 JSON 请求体
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dao.CreateUser(mysql.MysqlConnect, &models.User{
		UserName:     userInfo.UserName,
		UserAccount:  userInfo.UserAccount,
		UserPassword: userInfo.UserPassword,
		Status:       userInfo.Status,
		CreatedTime:  time.Now().Unix(),
		UpdatedTime:  time.Now().Unix(),
		IsDeleted:    userInfo.IsDeleted,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{
		"message": "User create successful",
		"user":    userInfo,
	})
}
