package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"golb/configs"
)

// jwtKey should be a `[]byte` type
var jwtKey = configs.JwtKey()

func keyFunc(tk *jwt.Token) (interface{}, error) {
	if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
	}
	return jwtKey, nil
}

// GenerateClaims generate claims
func GenerateClaims(slug string, id uint, roleID uint) *jwt.StandardClaims {
	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience:  slug,                             // aud 受众
		ExpiresAt: now.Add(time.Minute * 15).Unix(), // exp 失效时间
		Id:        Uint2String(id),                  // jti 编号 // TODO 使用无规律的ID避免重放攻击
		IssuedAt:  now.Unix(),                       // iat 签发时间
		Issuer:    "golb.sys",                       // iss 签发人
		NotBefore: now.Unix(),                       // nbf 生效时间
		Subject:   Uint2String(roleID),              // sub 主题
	}
	return claims
}

// GenerateTokenWithClaims generate a jwt token for the given claims
func GenerateTokenWithClaims(claims *jwt.StandardClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

// GenerateToken generate a token string
func GenerateToken(slug string, id uint, roleID uint) (token string, err error) {
	return GenerateTokenWithClaims(GenerateClaims(slug, id, roleID))
}

// JwtParse JwtParse
func JwtParse(tk string) (*jwt.Token, error) {
	return jwt.Parse(tk, keyFunc)
}

// RefreshToken RefreshToken
func RefreshToken(tk string) (string, error) {
	parseToken, err := JwtParse(tk)
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		newClaims := GenerateClaims(claims["aud"].(string), claims["jti"].(uint), claims["sub"].(uint))
		tokenStr, err := GenerateTokenWithClaims(newClaims)
		return tokenStr, err
	}
	return tk, err
}

// GetTokenRoleFromContext GetTokenRoleFromContext
func GetTokenRoleFromContext(ctx context.Context) (uint, error) {
	ginContext, ok := ctx.Value(ContextKey("GinContextKey")).(*gin.Context)
	if !ok {
		panic("could not retrive gin.Context")
	}
	token := ginContext.Request.Header.Get("Authorization")
	if len(token) == 0 {
		return 0, errors.New("Token incorrect")
	}
	parseToken, _ := JwtParse(token)
	if parseToken != nil && parseToken.Valid {
		claims, ok := parseToken.Claims.(jwt.MapClaims)
		if ok {
			return String2Uint(claims["sub"].(string)), nil
		}
	}
	return 0, errors.New("Token incorrect")
}

// CheckPermission Check the token in context whether has the correct role,
// return nil if everything ok;
// return an error if permission incorrect;
func CheckPermission(ctx context.Context, role uint) error {
	currentRole, err := GetTokenRoleFromContext(ctx)
	if err != nil {
		return errors.New("fail to resolve role from jwt in context")
	}
	if currentRole < role {
		return nil
	}
	return errors.New("permission deny")
}
