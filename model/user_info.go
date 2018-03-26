package model

import "gcoresys/common/mysql"

type UserInfo struct {
	mysql.BaseModel
	UserName string `json:"user_name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func GetDefaultUserInfo() UserInfo{
	return UserInfo{
		UserName:"测试",
		Age:18,
		Gender:"女",
	}
}