package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Computer(n))
}

func Computer(n int) string {
	var computer int
	var res string

	if n == 1 {
		res += fmt.Sprintf("%v компьютер", n)
	} else if n >= 2 && n <= 4 {
		res += fmt.Sprintf("%v компьютера", n)
	} else if n >= 5 && n <= 20 {
		res += fmt.Sprintf("%v компьютеров", n)
	} else if n > 20 {
		computer = n
		n %= 10
		if n == 1 {
			res += fmt.Sprintf("%v компьютер", computer)
		} else if n >= 2 && n <= 4 {
			res += fmt.Sprintf("%v компьютера", computer)
		} else {
			res += fmt.Sprintf("%v компьютеров", computer)
		}
	}
	return res
}
