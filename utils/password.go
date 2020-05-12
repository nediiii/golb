package utils

import (
	"encoding/hex"
	"log"

	"golang.org/x/crypto/bcrypt"

	"golb/configs"
)

var pepper = configs.Pepper()

func string2hex(s string) []byte {
	hex, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal("string2hex 失败", err)
	}
	return hex
}

func hex2string(h []byte) string {
	return hex.EncodeToString(h)
}

// IsCorrect check the password with `bcrypt` whether the it's valid
func IsCorrect(hashedPassword, plainPassword string) bool {
	return bcrypt.CompareHashAndPassword(string2hex(hashedPassword), []byte(plainPassword+pepper)) == nil
}

// Hash hash the password with `bcrypt` algorithm at cost 12
func Hash(password string) string {
	bcryptResult, err := bcrypt.GenerateFromPassword([]byte(password+pepper), 12)
	if err != nil {
		log.Fatal("hash encounter a error", err)
	}
	return hex2string(bcryptResult)
}
