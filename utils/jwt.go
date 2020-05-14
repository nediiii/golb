package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

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
func GenerateClaims(slug string, id uint) *jwt.StandardClaims {
	now := time.Now()
	claims := &jwt.StandardClaims{
		Audience:  slug,                             // aud 受众
		ExpiresAt: now.Add(time.Minute * 15).Unix(), // exp 失效时间
		Id:        Uint2String(id),                  // jti 编号
		IssuedAt:  now.Unix(),                       // iat 签发时间
		Issuer:    "golb.sys",                       // iss 签发人
		NotBefore: now.Unix(),                       // nbf 生效时间
		Subject:   "login",                          // sub 主题
	}
	return claims
}

// GenerateTokenWithClaims generate a jwt token for the given claims
func GenerateTokenWithClaims(claims *jwt.StandardClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

// GenerateToken generate a token string
func GenerateToken(slug string, id uint) (token string, err error) {
	return GenerateTokenWithClaims(GenerateClaims(slug, id))
}

// JwtParse JwtParse
func JwtParse(tk string) (*jwt.Token, error) {
	return jwt.Parse(tk, keyFunc)
}

// RefreshToken RefreshToken
func RefreshToken(tk string) (string, error) {
	parseToken, err := jwt.Parse(tk, keyFunc)
	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		newClaims := GenerateClaims(claims["aud"].(string), claims["jti"].(uint))
		tokenStr, err := GenerateTokenWithClaims(newClaims)
		return tokenStr, err
	}
	return tk, err
}
