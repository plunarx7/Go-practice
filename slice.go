package main

import (
	"fmt"
	"strconv"
	"sort"
)

func main() {
	slice := make([]int, 3)
	for {
		var input string
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)
		if input == "X" {
			return
		}
		if num, err := strconv.Atoi(input); err == nil {
			slice = append(slice, num)
			sort.Ints(slice)
			fmt.Println(slice)
		} else {
			fmt.Println("Enter numbers only!")
		}
	}
}
