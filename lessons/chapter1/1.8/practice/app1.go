package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	a1, _ := strconv.Atoi(arg1)
	a2, _ := strconv.Atoi(arg2)
	Signum(a1, a2)
}

func Signum(x int, y int) int {

	switch {
	case x > y:
		fmt.Println(x + y)
		return 1
	case x < y:
		fmt.Println(x - y)
		return 2
	default:
		return 0
	}
}
