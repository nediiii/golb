package utils

import (
	"strconv"
	"time"

	"log"
)

// Uint2String returns the string representation of i in the base 10.
func Uint2String(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// String2Uint returns the uint representation of s .
func String2Uint(s string) uint {
	i, err := strconv.Atoi(s)
	if err != nil || i < 0 {
		log.Println(i)
		log.Fatal("utils.String2Uint encounted an error!", err)
	}
	return uint(i)
}

// Time2UnixString Time2UnixString
func Time2UnixString(t *time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}
