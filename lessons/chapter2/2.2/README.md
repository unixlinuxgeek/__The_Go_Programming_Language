### 2.2 Объявления

В Go есть 4 основных объявления:

```var const type func```


```go
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Температура кипения = %g°F или %g°C\n", f,c)
	// Вывод
	//Температура кипения = 212°F или 100°С
}
```
Константа boilingF и функция main это объявление уровня пакета.
Каждая сущность уровня пакета видима для всех файлов уровня пакета.
Например, переменная boilingF или функция main

Локальные переменные f и c,
видимы только в пределах функции main

Пример с локальными константами freezingF и boilingF:
```go
// Ftoc выводит результаты преобразования двух температур
// по Фаренгейту в температуру по Цельсию.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0 // локальные переменные

	fmt.Printf("%g°F = %g°C", freezingF, boilingF)
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
```