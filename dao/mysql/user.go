package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go_web_app/model"
)

// md5 盐值
const secret = "salmonfishycooked"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	// 执行 SQL 语句
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

// InsertUser 向数据库中插入一条用户记录
func InsertUser(u *model.User) (err error) {
	// 密码加密
	password := encryptPassword(u.Password)

	// 执行 SQL 语句
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, u.UserID, u.Username, password)
	return
}

// encryptPassword 用于加密密码，返回一个十六进制的字符串
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
