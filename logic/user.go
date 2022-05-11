package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return
	}
	// 2.生成UID
	userID := snowflake.GetID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	err = mysql.InsertUser(user)
	return
}

func Login(p *models.ParamLogIn) (user *models.User, err error) {
	// 登录
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，所以在数据库中查找到的userID保存在user中
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Toekn = token
	return
}
