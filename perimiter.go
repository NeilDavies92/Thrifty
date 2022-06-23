package perimiter

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimiter(width, height float64) float64 {
	return 2 * (height + width)
}

func Area(width, height float64) float64 {
	return width * height
}
