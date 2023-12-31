// Возвращаем адрес локальной переменной.
// $ go run f.go
// 0xc000124000
// $ go run f.go
// 0xc0000180c0
// $ go run f.go
// 0xc0000b4000

package main

import (
	"fmt"
)

func main() {
	v := f()
	fmt.Println(v)
}

// Возвращаем адрес локальной переменной.
// Все вызовы возвращают разные значения.
func f() *int {
	v := 1
	// Локальная переменная v, будет существовать даже
	// после возврата из функции
	return &v // 1-ый запуск: 0xc000124000 ,2-ой запуск: 0xc0000180c0  3-ий запуск: 0xc0000b4000
}
