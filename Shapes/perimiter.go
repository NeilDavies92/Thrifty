package perimiter

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimiter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
