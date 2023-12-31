### Присваивание кортежу

 Прежде чем любая из переменных в левой части получит новое значение,
 вычисляются все выражения из правой части.
 
Например (рабочий пример):
```go
package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	// Сперва вычисляются правая часть err потом левая часть f   
	f, err := os.Open("dias.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
    }
	
	if _, err := io.WriteString(f, "Hi Dias Alan and Aidan!!!"); err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
		os.Exit(1)
    }
}
```

Еще примеры:
```go
v, ok =: m[key] // Поиск в отображении
v, ok := x.(T)  // Утверждение о типе
v, ok := <-ch   // Получение из канала  
```

```go
_, err = io.Copy(dst, src)  // Отбрасываем кол-во байтов
_, ok = x.(T)               // Проверка типа, игнорируем результат
```