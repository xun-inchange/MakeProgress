package util

import (
	"MakeProgress/gin-blog/gin-blog/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//jwt工具包

var jwtSecret = []byte(setting.JwtSecret)

//用户的状态和额外的元数据
type Claims struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

//生成token
func GenerateToken(userName, passWord string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //到期时间
	//这是负载的这一部分
	claims := Claims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			Issuer:    "li-xun",          //签发人
		},
	}
	//生成一个对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密生成token字符串
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

//解析token(根据tokenString获取到claims对象信息，进而获取到用户名与密码)
func ParseToken(tokenString string) (*Claims, error) {
	//用于解析鉴权的声明,方法内部主要是具体的解码和校验过程，最终返回*Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		//从token中获获取之前加密传入的claims struct对象并使用断言转换成我们的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		//这里还要判断一下token是否过期,如果没有过期我们就返回得到之前加密进去的Claims
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
