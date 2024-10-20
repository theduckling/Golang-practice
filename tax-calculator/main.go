package main

import "fmt"

func main() {
	var income int64
	var tax int64 = 0
	fmt.Scanf("%d", &income)
	if income > 100 {
		income -= 100
		tax += 5
		if income > 400 {
			income -= 400
			tax += 40
			if income-500 > 1 {
				income -= 500
				tax += 75
				if income > 1 {
					tax += income * 20 / 100
					fmt.Println(tax)
				}
			} else {
				tax += income * 15 / 100
				fmt.Println(tax)
			}
		} else {
			tax += income * 10 / 100
			fmt.Println(tax)
		}
	} else {
		tax += income * 5 / 100
		fmt.Println(tax)
	}
}
