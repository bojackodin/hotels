package main

import "fmt"

// from stdin
func main() {
	var min, max int
	fmt.Scan(&min, &max)
	s := []int{}

	for i := 0; i <= max; i++ {
		s = append(s, i)
	}

	s[1] = 0
	for i := 2; i <= max; i++ {
		if s[i] != 0 {
			for j := i + i; j <= max; j = j + i {
				s[j] = 0
			}
		}
	}

	res := []int{}
	if min == 0 {
		min = 1
	}
	for _, v := range s {
		if min <= v && v <= max {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}

// for testing
func SearchSimpleNumber(min, max int) []int {
	s := make([]int, 0, max)
	res := []int{}
	for i := 0; i <= max; i++ {
		s = append(s, i)
	}

	s[1] = 0
	for i := 2; i <= max; i++ {
		if s[i] != 0 {
			for j := i + i; j <= max; j = j + i {
				s[j] = 0
			}
		}
	}

	if min == 0 {
		min = 1
	}
	for _, v := range s {
		if min <= v && v <= max {
			res = append(res, v)
		}
	}

	return res
}
