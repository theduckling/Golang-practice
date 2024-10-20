package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	Key   string
	Value string
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// nStr := scanner.Text()
	// nStr = strings.TrimSpace(nStr)
	books := make(map[string]string)
	n, _ := strconv.Atoi(scanner.Text())
	// if err != nil {
	// 	fmt.Println("error")
	// 	return
	// } else if n < 1 || n > 100 {
	// 	fmt.Println("error")
	// 	return
	// }
	for i := 0; i < n; i++ {
		scanner.Scan()
		scores := strings.Fields(scanner.Text())
		ISBN := scores[1]
		// if err != nil {
		// 	fmt.Println("error")
		// 	return
		// }
		switch {
		case scores[0] == "ADD":
			title := strings.Join(scores[2:], " ")

			books[ISBN] = title
		case scores[0] == "REMOVE":
			for key := range books {
				if key == ISBN {
					delete(books, key)
				}
			}
		}
	}
	var sortedSlice []kv
	for k, v := range books {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		if sortedSlice[i].Value == sortedSlice[j].Value {
			first, _ := strconv.Atoi(sortedSlice[i].Key)
			second, _ := strconv.Atoi(sortedSlice[j].Key)
			return first < second
		}
		return sortedSlice[i].Value < sortedSlice[j].Value // Sort by value
	})
	// fmt.Println("Sorted map by values:")
	for _, kv := range sortedSlice {
		fmt.Printf("%s \n", kv.Key)
	}

}
