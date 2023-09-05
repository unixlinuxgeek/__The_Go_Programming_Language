package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Сперва вычисляются правая часть err потом левая часть f
	f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
	}

	// n это количество записанных байтов
	if n, err := io.WriteString(f, "Hi Dias Alan and Aidan!!!\n"); err == nil {
		fmt.Println(n)
	} else {
		fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
	}
}
