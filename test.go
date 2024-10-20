// package main

// import (
// 	// "buio"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	// num := 10

// 	// if num < 100 {
// 	//     num := 5 //shadow variable
// 	//     fmt.Printf("num value inside if block: %d\n", num)
// 	// }

// 	// fmt.Printf("num value outside if block: %d", num)

// 	//--------------------------------------------------------------------
// 	// scanner := bufio.NewScanner(os.Stdin)
// 	// scanner.Scan()
// 	// words := strings.Fields(scanner.Text())
// 	// fmt.Print(words)

// 	// var n int = int(fmt.Scanf("%v", nil))
// 	// print(n)
// 	//--------------------------------------------------------------------
// 	//--------------------------------------------------------------------

// 	var players = map[string]int{
// 		"Ronaldo": 7,
// 		"Messi":   10,
// 		"Neymar":  11,
// 	}

// 	for key, value := range players {
// 		fmt.Printf("Key : %v, Value : %v\n", key, value)
// 	}

// }
// ______________________________________________________________________
//higher order function
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("%e", 33)
// }

//	func applyFunction(f func(int) int, num int) int {
//	    return f(num)
//	}
//
//	func square(n int) int {
//	    return n * n
//	}
//
//	func cube(n int) int {
//	    return n * n * n
//	}
//
//	func main() {
//	    fmt.Println(applyFunction(square, 5))
//	    fmt.Println(applyFunction(cube, 5))
package main

import (
	"fmt"
	"math"
)

func main() {
	x := 12
	count := 0
	var reverse int
	number := x
	num := x
	for number != 0 {
		number /= 10
		count += 1
	}
	for i := 0; i < count; i++ {
		reverse = reverse + (num%10)*int(math.Pow(10, float64(count-i-1)))
		num = num / 10
	}
	fmt.Print(reverse, "\n", count)

}
