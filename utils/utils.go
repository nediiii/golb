package utils

import "strconv"

// Uint2string returns the string representation of i in the base 10.
func Uint2string(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}
