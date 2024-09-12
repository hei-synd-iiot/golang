package toto

type Shape interface { // (1)!
	SetHeight(height int) // (2)!
	SetWidth(width int)
	CalculateArea() int
}

type rectangle struct { // (3)!
	height, width int
}

func NewRectangle(h, w int) Shape { // (4)!
	return &rectangle{height: h, width: w}
}

func (r *rectangle) SetHeight(height int) { // (5)!
	r.height = height
}

func (r *rectangle) SetWidth(width int) {
	r.width = width
}

func (r *rectangle) CalculateArea() int {
	return r.height * r.width
}
