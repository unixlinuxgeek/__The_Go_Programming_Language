// Ftoc выводит результаты преобразования двух температур
// по Фаренгейту в температуру по Цельсию.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0 // локальные переменные

	fmt.Printf("%g°F = %g°C\n", freezingF, fToc(freezingF))
	// 32°F = 0°C
	fmt.Printf("%g°F = %g°C\n", boilingF, fToc(boilingF))
	// 212°F = 100°C
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
