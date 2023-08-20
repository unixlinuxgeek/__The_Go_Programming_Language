// Измените программу fetch так, чтобы к каждому аргументу URL
// автоматически добавлялся префикс ```http://``` в случае отсутствия в нем такового.
// Можете воспользоваться функцией ```strings.HasPrefix```.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

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
		fmt.Printf("%s", b)
	}
}
