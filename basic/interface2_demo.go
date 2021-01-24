package basic

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

type Cube struct {
	Side float64
}

func (c Cube) Area() float64 {
	return 6* (c.Side*c.Side)
}

func (c Cube) Volume() float64 {
	return c.Side*c.Side*c.Side
}


