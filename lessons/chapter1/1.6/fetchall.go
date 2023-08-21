// Fetchall выполняет параллельную выбоорку URL и сообщает
// о затраченном времени и размере ответа для каждого из них.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan []byte)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		// Получение из канала ch
	}
	for range os.Args[1:] {
		f, err := os.Create("1.txt")
		defer f.Close()
		if err != nil {
			ch <- []byte("Error")
		}

		f.Write(<-ch)
		defer f.Close()

		fmt.Println(<-ch)
		// Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- []byte) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- []byte(err.Error()) // Отправка в канал ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // Исключение утечки ресурсов
	if err != nil {
		//ch <- fmt.Sprintf("while reading %s: %v", url, err)
		ch <- []byte(fmt.Sprintf("while reading %s: %v", url, err))
		return
	}
	secs := time.Since(start).Seconds()
	ch <- []byte(fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url))
}
