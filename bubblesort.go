package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	// swap the value of i with i+1
	if index < 0 || index >= len(slice)-1 {
		fmt.Println("Index out of range")
		return
	}
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func BubbleSort(slice []int) {
	var swapped bool
	for i := 0; i < len(slice)-1; i++ {
		swapped = false
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func ReadSlice() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a list of numbers separated by spaces: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strSlice := strings.Split(input, " ")
	intSlice := make([]int, 0, len(strSlice))
	for _, strNum := range strSlice {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println("Error converting input to integer: ", err)
			return nil
		}
		intSlice = append(intSlice, num)
	}
	fmt.Println("Slice of integers: ", intSlice)
	return intSlice
}

func main() {
	sliceNum := ReadSlice()
	BubbleSort(sliceNum)
	fmt.Println("Sorted slice: ", sliceNum)
}
