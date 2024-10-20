package arsalan

import (
	"fmt"
	"math"
)

func arsalan() {
	var tax float64 = 0
	var income int = 0
	fmt.Scanf("%d", &income)

	if income <= 10_000 {
		// 5
		tax += float64(income) * 0.05
	} else if (income > 10_000) && (income <= 50_000) {
		// 5
		tax += 10_000 * 0.05
		income -= 10_000
		// 10
		tax += float64(income) * 0.1
	} else if (income > 50_000) && (income <= 100_000) {
		// 5
		tax += 10_000 * 0.05
		income -= 10_000
		// 10
		tax += 40_000 * 0.1
		income -= 40_000
		// 15
		tax += float64(income) * 0.15
	} else if income > 100_000 {
		// 5
		tax += 10_000 * 0.05
		income -= 10_000
		// 10
		tax += 40_000 * 0.1
		income -= 40_000
		// 15
		tax += 50_000 * 0.15
		income -= 50_000
		// 20
		tax += float64(income) * 0.2
	}

	fmt.Println(math.Floor(tax))
}
