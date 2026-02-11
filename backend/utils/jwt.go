package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims JWT令牌声明结构体
// 包含用户自定义声明和JWT标准声明
type Claims struct {
	UserID string `json:"userId"` // 用户ID
	Role   string `json:"role"`   // 用户角色（manufacturer/automaker/aftersale）
	jwt.RegisteredClaims          // JWT标准声明
}

// GenerateToken 生成JWT令牌
// 参数：
//   - userID: 用户ID
//   - role: 用户角色
//   - secret: JWT签名密钥
//   - expireHours: 令牌过期时间（小时）
// 返回：
//   - string: 生成的JWT令牌
//   - error: 生成过程中的错误
func GenerateToken(userID string, role string, secret string, expireHours int) (string, error) {
	// 设置令牌过期时间
	expireAt := time.Now().Add(time.Duration(expireHours) * time.Hour)
	
	// 创建自定义声明
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
		},
	}
	
	// 创建令牌（使用HS256算法签名）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// 使用密钥签名令牌并返回字符串
	return token.SignedString([]byte(secret))
}

// ParseToken 解析并验证JWT令牌
// 参数：
//   - tokenString: JWT令牌字符串
//   - secret: JWT签名密钥
// 返回：
//   - *Claims: 令牌声明（包含用户信息）
//   - error: 解析或验证过程中的错误
func ParseToken(tokenString string, secret string) (*Claims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为HS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		// 返回密钥用于验证签名
		return []byte(secret), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// 验证令牌有效性并提取声明
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}
