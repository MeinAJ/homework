package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义 Claims 结构体
type CustomClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTConfig JWT 配置
type JWTConfig struct {
	SigningKey  []byte
	ExpiresTime time.Duration
	Issuer      string
}

var (
	jwtConfig = JWTConfig{
		SigningKey:  []byte("VI1Er5p6jkA57ZK94BMJ5rYpIs6lGW4ydMMqgSf6cyg="),
		ExpiresTime: time.Hour * 24,
		Issuer:      "blockchain",
	}
)

// GenerateToken 生成 JWT Token
func GenerateToken(userID uint, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.ExpiresTime)),
			Issuer:    jwtConfig.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtConfig.SigningKey)
}

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "请求头中未找到Authorization信息"})
			return
		}

		// 校验 Token 格式
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Token格式不正确"})
			return
		}

		// 解析并验证 Token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtConfig.SigningKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Token验证失败: " + err.Error()})
			return
		}

		// 类型断言获取 Claims
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// 将用户信息存入上下文
			c.Set("userID", claims.UserID)
			c.Set("username", claims.Username)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的Token声明"})
		}
	}
}
