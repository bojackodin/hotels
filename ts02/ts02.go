package main

import "fmt"

func main() {
	// from stdin
	var n, input int
	fmt.Scanln(&n)
	s := make([]int, 0, n)
	res := []int{}
	max := 0
	counter := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		if max < input {
			max = input
		}
		s = append(s, input)
	}

	for i := 2; i <= 6; i++ {
		for _, v := range s {
			if v%i == 0 {
				counter++
			} else {
				counter = 0
				break
			}
		}
		if counter == 3 {
			res = append(res, i)
			counter = 0
		}
	}
	fmt.Println(res)
}

// for testing
func NOD(s []int) []int {
	res := []int{}
	max := 0
	counter := 0
	for i := 0; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	n := len(s)
	for i := 2; i <= max; i++ {
		for _, v := range s {
			if v%i == 0 {
				counter++
			} else {
				counter = 0
				break
			}
		}
		if counter == n {
			res = append(res, i)
			counter = 0
		}
	}
	return res
}
