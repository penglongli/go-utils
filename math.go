package goutils

import "math"

// Round keep n bits decimals
func Round(f float64, n int) float64 {
	if n < 0 {
		return 0
	}
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
