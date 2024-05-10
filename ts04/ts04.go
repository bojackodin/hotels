package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// TODO rectify for numbers more 10
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number 1 to 10: ")
	input, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}
	if num > 10 {
		fmt.Println("Input more than 10. Please enter a valid number.")
		return
	}
	width := 1
	maxResult := num * num
	for maxResult > 0 {
		maxResult /= 10
		width++
	}
	if num >= 10 {
		fmt.Print("  ")
	} else {
		fmt.Print(" ")
	}
	for i := 1; i <= num; i++ {
		if i >= 10 && i < 100 {
			spacePrint(width - 1)
		} else if i >= 100 && i < 1000 {
			spacePrint(width - 2)
		} else {
			spacePrint(width)
		}

		fmt.Printf("%d", i)
	}

	fmt.Println()
	for i := 1; i <= num; i++ {
		if num >= 10 {
			if i < 10 {
				spacePrint(1)
				fmt.Printf("%d", i)
			} else {
				fmt.Printf("%d", i)
			}
		} else {
			fmt.Printf("%d", i)
		}
		for j := 1; j <= num; j++ {
			if i*j >= 10 && i*j < 100 {
				spacePrint(width - 1)
			} else if i*j >= 100 && i*j < 1000 {
				spacePrint(width - 2)
			} else {
				spacePrint(width)
			}
			fmt.Printf("%d", i*j)
		}
		fmt.Println()
	}
}

func spacePrint(width int) {
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}
}
