package tempconv

import "fmt"

type Cel float64
type Far float64
type Kel float64

const (
	AbsoluteZeroC Cel = -273.15
	FreezingC     Cel = 0
	BoilingC      Cel = 100
	KelvinZeroC   Cel = -273.15
)

func (c Cel) String() string {
	return fmt.Sprintf("%g ℃", c)
}

func (f Far) String() string {
	return fmt.Sprintf("%g ℉", f)
}

func (k Kel) String() string {
	return fmt.Sprintf("%g °K", k)
}
