package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return errors.New("error")
	}
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) < i-1 {
		return "", errors.New("error")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) < i-1 {
		panic(errors.New("error"))
	}
	s1 := b.shapes[:i]
	s2 := b.shapes[i+1:]
	s3 := []Shape{}
	s3 = append(s3, s1...)
	s3 = append(s3, s2...)
	b.shapes = s3
	return b.shapes[i], nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) < i-1 {
		panic(errors.New("error"))
	}
	s1 := b.shapes[:i]
	s2 := b.shapes[i+1:]
	s3 := []Shape{}
	s3 = append(s3, s1...)
	s3 = append(s3, shape)
	s3 = append(s3, s2...)
	b.shapes = s3

	return b.shapes[i], nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	p := []float64{}
	var sum float64
	for _, v := range b.shapes {
		switch s := v.(type) {
		case Circle:
			p = append(p, (2*3.14)*s.Radius)
		case Triangle:
			p = append(p, 3*s.Side)
		case Rectangle:
			p = append(p, (s.Height+s.Weight)*2)
		}
	}
	for _, v2 := range p {
		sum += v2
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	p := []float64{}
	var sum float64
	for _, v := range b.shapes {
		switch s := v.(type) {
		case Circle:
			p = append(p, (3.14*s.Radius)*(3.14*s.Radius))
		case Triangle:
			p = append(p, (s.Side*s.Side)/2)
		case Rectangle:
			p = append(p, s.Height*s.Weight)
		}
	}
	for _, v2 := range p {
		sum += v2
	}
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	c := 0
	for i, v := range b.shapes {
		switch v.(type) {
		case Circle:
			c++
			s1 := b.shapes[:i]
			s2 := b.shapes[i+1:]
			s3 := []Shape{}
			s3 = append(s3, s1...)
			s3 = append(s3, s2...)
			b.shapes = s3
		}
	}
	if c <= 0 {
		return errors.New("error")
	}
	return errors.New("error")
}
