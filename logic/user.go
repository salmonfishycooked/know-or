package logic

import (
	"go_web_app/dao/mysql"
	"go_web_app/model"
	"go_web_app/pkg/snowflake"
)

func SignUp(p *model.ParamSignUp) (err error) {
	// 判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	// 构造用户实例
	userID := snowflake.GenID() // 生成 UID
	u := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 插入数据库
	return mysql.InsertUser(u)
}
