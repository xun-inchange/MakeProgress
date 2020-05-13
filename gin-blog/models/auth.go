package models

//登录模型

type Auth struct {
	ID       int    `gorm:"primary key" json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

//检查用户是否存在
func CheckAuth(userName, passWord string) bool {
	var auth Auth
	db.Select("id").Where(Auth{PassWord: passWord, UserName: userName}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
