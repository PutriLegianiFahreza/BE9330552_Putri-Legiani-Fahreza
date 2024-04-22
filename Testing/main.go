package Testing

import (
	"fmt"
)

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func main() {
	t := Triangle{
		Base:   10.0,
		Height: 5.0,
	}
	fmt.Printf("luas segitiga adalah : %.2f\n", t.Area())
}
