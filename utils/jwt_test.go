package utils

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

// make sure jwt can work
func TestJWT(t *testing.T) {
	var testCase = []struct {
		slug   string
		id     uint
		roleID uint
	}{
		{"slug", 1, 2},
		{"", 0, 0},
	}

	for _, tC := range testCase {
		if token, err := GenerateToken(tC.slug, tC.id, tC.roleID); err != nil {
			t.Error("GenerateToken wrong")
			newToken, err := RefreshToken(token)
			if err != nil {
				t.Error("RefreshToken wrong")
			}
			parseToken, err := JwtParse(newToken)
			if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
				if tC.slug != claims["aud"].(string) {
					t.Error("RefreshToken has a mistakeon slug")
				}
				if tC.id != claims["jti"].(uint) {
					t.Error("RefreshToken has a mistake on id")
				}
				if tC.roleID != claims["sub"].(uint) {
					t.Error("RefreshToken has a mistake on roleID")
				}
			}
		}
	}

}
