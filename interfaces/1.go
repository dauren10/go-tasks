// Объявление интерфейса
type Shape interface {
	Area() float64
}

// Реализация интерфейса для типа Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Реализация интерфейса для типа Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	// Использование интерфейса для обработки разных типов данных
	shapes := []Shape{Circle{Radius: 2.0}, Rectangle{Width: 3.0, Height: 4.0}}

	for _, shape := range shapes {
		fmt.Println("Площадь:", shape.Area())
	}
}
