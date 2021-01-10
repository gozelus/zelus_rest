package core

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	UserID int64
	*jwt.StandardClaims
}

var day int64 = 24 * 60 * 60
var claims = CustomClaims{
	StandardClaims: &jwt.StandardClaims{
		ExpiresAt: 7 * day,
	},
}

func ValidateToken(key, tokenStr string) (int64, string, error) {
	var err error
	var jwtToken *jwt.Token
	if jwtToken, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return key, nil
	}); err != nil {
		return 0, "", err
	}
	if cc, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
		// 看下 token 是否快过期了，如果快过期了，就要生成一个新的给客户端使用
		s := jwtToken.Claims.(*CustomClaims)
		if s.ExpiresAt-int64(time.Now().Unix()) < 3*day {
			newTokenStr, err := NewToken(key)
			if err != nil {
				return 0, "", err
			}
			return cc.UserID, newTokenStr, nil
		}
	}
	return 0, "", err
}
func NewToken(key string) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenStr, err := newToken.SignedString(key)
	if err != nil {
		return "", err
	}
	return newTokenStr, nil
}
