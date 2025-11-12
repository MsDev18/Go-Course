package mathematic

import (
	"math"
)

func PowerH(base, power float64) float64 {
	sum := float64(1)
	for i := float64(0); i < power; i++ {
		sum *= base
	}
	return sum
}

func PowerMath(base, power float64) float64 {
	return math.Pow(base, power)
}


func PowerSlc (base , power float64) float64 {
	result := make([]float64, int(power))

	sum := float64(1)
	for i,_ := range result {
		sum *= base + float64(i)
	}
	return sum
}