package main

import (
	"math"
)

func reverse(x int) int {
	var numbers []int
	var flag bool
	if x < 0 {
		x = -x
		flag = true
	}
	for x > 0 {
		num := x % 10
		numbers = append(numbers, num)
		x = x / 10
	}
	length := len(numbers)
	var result int
	for i := 0; i < length; i++ {
		base := int(math.Pow(10, float64(length-i-1)))
		result = result + numbers[i]*base
	}

	if flag {
		result = -result
	}
	if (result > int(math.Pow(2, 31)-1)) || (result < -(int(math.Pow(2, 31)))) {
		return 0
	}

	return result
}

func main() {
	reverse(123)
	reverse(-123)
	reverse(120)
}
