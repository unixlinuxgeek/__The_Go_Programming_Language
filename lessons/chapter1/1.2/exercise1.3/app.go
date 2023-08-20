// Поэксперементируйте с измерением разницы времени выполнения
// потенциально неэффективных версии и версии с применением strings.Join.

// В разделе exercise1.5.6 демонстрируются часть пакета time.
// А в разделе 11.4 - как написать тест производительности для ее систематической оценки.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	//for i := exercise1.5; i < len(os.Args); i++ {
	//	s += strconv.Itoa(i) + " \n" + sep + os.Args[i]
	//
	//}

	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += strings.Join(os.Args, sep) + " \n" // 700-1000
		//s += strconv.Itoa(i) + " \n" + sep + os.Args[i] // 400-550
	}

	m := time.Since(start).Nanoseconds()
	fmt.Println(m)
	fmt.Println("------------------------------------------------------")
}
