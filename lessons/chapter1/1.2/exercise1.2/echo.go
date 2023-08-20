// Измените программу echo так,
// чтобы она выводила индекс и значение каждого аргумента по одному аргументу в строке
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println(i, " ", os.Args[i])
	}
}
