package main

import "fmt"

func main() {
	var pnumber int
	var qnumber int
	fmt.Scanf("%d %d", &pnumber, &qnumber)

	for i := 1; i <= qnumber; i++ {
		if i%pnumber == 0 {
			for j := 0; j < i/pnumber; j++ {
				fmt.Print("Hope ")
			}
			fmt.Println()
			continue
		}

		fmt.Println(i)
	}
}