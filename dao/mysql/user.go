package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"go_web_app/model"
	"go_web_app/pkg/e"
)

// md5 盐值
const secret = "salmonfishycooked"

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	// 执行 SQL 语句
	sqlStr := `SELECT count(user_id) FROM user WHERE username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		// 查询数据库失败
		return err
	}
	if count > 0 {
		return e.ErrorUserExist
	}
	return
}

// InsertUser 向数据库中插入一条用户记录
func InsertUser(u *model.User) (err error) {
	// 密码加密
	password := encryptPassword(u.Password)

	// 执行 SQL 语句
	sqlStr := `INSERT INTO user(user_id, username, password) VALUES(?, ?, ?)`
	_, err = db.Exec(sqlStr, u.UserID, u.Username, password)
	return
}

// encryptPassword 用于加密密码，返回一个十六进制的字符串
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(u *model.User) (err error) {
	oPassword := u.Password // 记录用户输入的密码

	// 执行 SQL 语句
	sqlStr := `SELECT user_id, username, password FROM user WHERE username = ?`
	err = db.Get(u, sqlStr, u.Username)
	if err == sql.ErrNoRows {
		return e.ErrorUserNotExist
	}
	if err != nil {
		// 查询数据库失败
		return
	}

	// 判断密码是否一致
	password := encryptPassword(oPassword)
	if password != u.Password {
		return e.ErrorInvalidPassword
	}
	return
}
