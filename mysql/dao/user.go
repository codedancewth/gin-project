package dao

import (
	"gin-project/mysql/models"
	"gorm.io/gorm"
	"log"
)

// GetUserList 获取用户的基础数据列表
func GetUserList(db *gorm.DB) (err error, users []*models.User) {
	err = db.Table((&models.User{}).Table()).Find(&users).Error
	if err != nil {
		log.Fatalf("find user list error [%v]", err)
	}
	return err, users
}

// GetUserListById 获取用户的基础数据列表
// param: ID
func GetUserListById(db *gorm.DB, id int64) (err error, users []*models.User) {
	err = db.Table((&models.User{}).Table()).Where("id = ?", id).Find(&users).Error
	if err != nil {
		log.Fatalf("find user list error [%v]", err)
	}
	return err, users
}

// CreateUser 创建用户的数据
func CreateUser(db *gorm.DB, user *models.User) (err error) {
	if err = db.Table((&models.User{}).Table()).Create(user).Error; err != nil {
		log.Printf("save user error [%v]", err)
	}

	return err
}

// SaveUser 存储用户的数据
func SaveUser(db *gorm.DB, user *models.User) {
	if err := db.Table((&models.User{}).Table()).Save(user).Error; err != nil {
		log.Fatalf("save user error [%v]", err)
	}
}
