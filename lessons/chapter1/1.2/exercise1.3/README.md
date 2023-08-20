### 1.3

Поэксперементируйте с измерением разницы времени выполнения
потенциально неэффективных версии и версии с применением strings.Join.

В разделе 1.6 демонстрируются часть пакета time.
А в разделе 11.4 - как написать тест производительности для ее систематической оценки.

Strings.Join: 700-1000
Конкатенация строк: 400-550

Конкатенация строк работает быстрее Strings.Join

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var s, sep string

	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		//s += strings.Join(os.Args, sep) + " \n" // 700-1000
		s += strconv.Itoa(i) + " \n" + sep + os.Args[i] // 400-550
	}

	m := time.Since(start).Nanoseconds()
	fmt.Println(m)
	fmt.Println("------------------------------------------------------")
}
```