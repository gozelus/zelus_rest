package core

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtUtils struct {
	Key    string
	Secret string
}

type CustomClaims struct {
	UserID int64
	*jwt.StandardClaims
}

var day int64 = 24 * 60 * 60

func NewJwtUtils(key, secret string) *JwtUtils {
	return &JwtUtils{
		Key:    key,
		Secret: secret,
	}
}

func (u *JwtUtils) ValidateToken(tokenStr string) (int64, string, error) {
	var err error
	var jwtToken *jwt.Token
	if jwtToken, err = jwt.ParseWithClaims(tokenStr, &CustomClaims{
		StandardClaims: &jwt.StandardClaims{},
	}, func(token *jwt.Token) (i interface{}, e error) {
		return u.Key, nil
	}); err != nil {
		return 0, "", err
	}
	if cc, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
		// 看下 token 是否快过期了，如果快过期了，就要生成一个新的给客户端使用
		s := jwtToken.Claims.(*CustomClaims)
		if s.ExpiresAt-int64(time.Now().Unix()) < 3*day {
			newTokenStr, err := u.NewToken(cc.UserID)
			if err != nil {
				return 0, "", err
			}
			return cc.UserID, newTokenStr, nil
		}
	}
	return 0, "", err
}
func (u *JwtUtils) NewToken(uid int64) (string, error) {
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		StandardClaims: &jwt.StandardClaims{},
		UserID:         uid,
	})
	newTokenStr, err := newToken.SignedString(u.Secret)
	if err != nil {
		return "", err
	}
	return newTokenStr, nil
}
