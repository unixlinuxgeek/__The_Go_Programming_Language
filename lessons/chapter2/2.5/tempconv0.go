// Пакет tempconv выполняет вычисления температур
// по Цельсию и по Фаренгейту

package tempconv

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CTof(c Celsius) Fahrenheit {
	// Fahrenheit(c*9/5 + 32) это преобразование типа, а не вызов функций
	return Fahrenheit(c*9/5 + 32)
}

func FToc(f Fahrenheit) Celsius {
	// Celsius((f - 32) * 5 / 9) это преобразование типа, а не вызов функций
	return Celsius((f - 32) * 5 / 9)
}
