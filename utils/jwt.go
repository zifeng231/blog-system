package utils

import (
	"blog-system/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	User models.User `json:"user"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			Issuer:    "blogsystem",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET)
}

// ParseToken 解析并验证 JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// 使用 jwt.ParseWithClaims 解析 token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{}, // 创建一个空的 Claims 实例，用于接收解析后的数据
		func(token *jwt.Token) (interface{}, error) {
			// 返回签名使用的密钥
			return JWT_SECRET, nil
		},
	)
	// 检查解析结果
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// token 有效，返回 claims
		return claims, nil
	} else {
		// token 无效或解析失败
		return nil, err
	}
}
