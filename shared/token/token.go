package token

import (
	"coderhub/conf"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

// 定义一个密钥，用于签名和验证 Token
var jwtKey = []byte(conf.JWTSecret)

func GenerateAuthorization(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		ID:        strconv.FormatInt(userID, 10),                             // 唯一标识
		Issuer:    "",                                                        // 签发者
		IssuedAt:  jwt.NewNumericDate(time.Now()),                            // 签发时间
		Subject:   "",                                                        // 主题ß
		Audience:  jwt.ClaimStrings{"", ""},                                  // 接收者
		NotBefore: jwt.NewNumericDate(time.Now()),                            // 生效时间
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 60 * time.Minute)), // 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
