// Измените программу fetch так, чтобы она выводила код состояния HTTP,
// содержащиеся в ```resp.Status```.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Printf("%s\n", resp.Status)
	}
}
