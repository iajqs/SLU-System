/**
* Created by cks
* Date: 2020-12-05
* Time: 15:35
*/

package dao

// import (
// 	"SLU-System/db" 
// 	"github.com/pkg/errors"
// 	"time"
// )

// var dbIns = db.GetDb("SLU-System")

// type User struct {
// 	Id         int `gorm:"primary_key"`
// 	UserName   string
// 	Password   string
// 	CreateTime time.Time
// 	db.DbSLU
// }

// func (u *User) TableName() string {
// 	return "user"
// }

// func (u *User) Add() (userId int, err error) {
// 	if u.UserName == "" || u.Password == "" {
// 		return 0, errors.New("user_name or password empty!")
// 	}
// 	oUser := u.CheckHaveUserName(u.UserName)
// 	if oUser.Id > 0 {
// 		return oUser.Id, nil
// 	}
// 	u.CreateTime = time.Now()
// 	if err = dbIns.Table(u.TableName()).Create(&u).Error; err != nil {
// 		return 0, err
// 	}
// 	return u.Id, nil
// }

// func (u *User) CheckHaveUserName(userName string) (data User) {
// 	dbIns.Table(u.TableName()).Where("user_name=?", userName).First(&data)
// 	return
// }

// func (u *User) GetUserNameByUserId(userId int) (userName string) {
// 	var data User
// 	dbIns.Table(u.TableName()).Where("user_id=?", userId).First(&data)
// 	return data.UserName
// }
