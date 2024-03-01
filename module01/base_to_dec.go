package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	charset := "0123456789ABCDEF"
	cpt := len(value) - 1
	sum := 0
	for _, val := range value {
		num := strings.Index(charset, string(val))
		sum += num * int(math.Pow(float64(base), float64(cpt)))
		cpt--
	}
	return sum
}
