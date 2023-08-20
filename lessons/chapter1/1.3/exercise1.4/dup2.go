// Упражнение exercise1.5.4
// Измените программу dup2 так, чтобы она выводила имена всех файлов,
// в которых найдены повторяющиеся строки.

// Ищем комментарий 'Решение задачи exercise1.5.4 !!!'

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLinesN(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			countLinesN(f, counts)
			f.Close()
		}
		// Решение задачи exercise1.5.4 !!!
		for s, i := range counts {
			if i > 1 {
				fmt.Println(i, s)
			}
		}
	}
}

func countLinesN(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
