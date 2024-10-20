package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Invalid number of countries")
		return
	}

	countries := make(map[string]string)
	for i := 0; i < n; i++ {
		country_num, _ := reader.ReadString('\n')
		countries[strings.Fields(country_num)[1]] = strings.Fields(country_num)[0]

	}

	n_numberstr, _ := reader.ReadString('\n')
	n_numberstr = strings.TrimSpace(n_numberstr)
	n_number, err := strconv.Atoi(n_numberstr)
	// fmt.Printf("%d", n_number)
	numbers := make([]string, n_number)

	for i := 0; i < n_number; i++ {
		number, _ := reader.ReadString('\n')
		number = strings.TrimSpace(number)
		// numbers = append(numbers, number)
		numbers[i] = number

	}
	for _, num := range numbers {
		val, ok := countries[num[:3]]
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("Invalid Number")
		}
	}
}
