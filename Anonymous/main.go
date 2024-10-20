package main

import (
	"math"
)

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	min := math.MaxInt64
	if len(*numbers) == 0 {
		return 0
	} else {
		for _, value := range *numbers {
			if value < min {
				min = value
			}
		}
		return min
	}
}

func ReverseSlice(numbers *[]int) {
	Rnumbers := *numbers
	Nnumbers := len(*numbers)
	for i := 0; i < Nnumbers/2; i++ {
		Rnumbers[Nnumbers-i-1], Rnumbers[i] = Rnumbers[i], Rnumbers[Nnumbers-i-1]
	}
}

func SwapElements(numbers *[]int, i, j int) {
	Rnumbers := *numbers
	if (i < len(Rnumbers) && i >= 0) && (j < len(Rnumbers) && j >= 0) {

		Rnumbers[i], Rnumbers[j] = Rnumbers[j], Rnumbers[i]
	} else {
		return
	}
}
