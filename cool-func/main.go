package main

import (
	"math"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	//TODO
	if x < 0 {
		return false
	} else {
		sqrt := int(math.Sqrt(float64(x)))
		return sqrt*sqrt == x
	}

}

func IsPalindrome(x int) bool {
	count := 0
	var reverse int
	number := x
	for number != 0 {
		number /= 10
		count += 1
	}
	number = x
	for i := 0; i < count; i++ {
		reverse = reverse + (number%10)*int(math.Pow(10, float64(count-i-1)))
		number = number / 10
	}
	return reverse == x
}

//------------------------------------------------------

func Abs(num int) int {
	if num < 0 {
		num = num * -1
	}
	return num
}

func Cube(num int) int {
	//TODO
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	result := make([]int, 0)
	for i, element := range input {
		if f(element) {
			result = append(result, input[i])
		}
	}

	return result
}

func Map(input []int, m MapperFunc) []int {
	for i, element := range input {
		input[i] = m(element)
	}
	return input
}
