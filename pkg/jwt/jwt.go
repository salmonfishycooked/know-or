package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go_web_app/pkg/e"
	"time"
)

var ErrorInvalidToken = errors.New(e.CodeInvalidToken.Msg())

// TokenExpireDuration 是 token 的过期时间
const TokenExpireDuration = time.Hour * 2

// Issuer 指 token 签发人
const Issuer = "salmonfishycooked"

// mySecret 是签名盐值
var mySecret = []byte("salmonfishycooked.good@me!!!")

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt 包自带的 jwt.StandardClaims 只包含了官方字段
// 我们这里需要额外记录一些字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 用来生成 JWT，会返回一个生成好的 token
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己声明的数据
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    Issuer,                                     // 签发人
		},
	}
	// 使用指定的加密方法生成待签名的对象
	obj := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的 secret 签名并获得 Base64 编码后的字符串 token
	return obj.SignedString(mySecret)
}

// ParseToken 用来解析 token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	mc := &MyClaims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		mc,
		func(token *jwt.Token) (interface{}, error) {
			return mySecret, nil
		},
	)
	if err != nil {
		return nil, err
	}
	// 检验 token 是否有效
	if token.Valid {
		return mc, nil
	}
	return nil, ErrorInvalidToken
}
