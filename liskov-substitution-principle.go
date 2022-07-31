package main

import (
	"fmt"
)

type Area interface {
	GetArea() float64
}

type AreaQuadrado interface {
	Area
	GetSize() float64
	SetSize(size float64)
}

type AreaRetangulo interface {
	Area
	GetWidth() float64
	GetHeight() float64
	SetWidth(width float64)
	SetHeight(height float64)
}

type Quadrado struct {
	size float64
}

type Retangulo struct {
	width  float64
	height float64
}

func (q *Quadrado) GetArea() float64 {
	return q.size * q.size
}

func (r *Retangulo) GetArea() float64 {
	return r.width * r.height
}

func (q *Quadrado) GetSize() float64 {
	return q.size
}

func (q *Quadrado) SetSize(size float64) {
	q.size = size
}

func (r *Retangulo) GetWidth() float64 {
	return r.width
}

func (r *Retangulo) GetHeight() float64 {
	return r.height
}

func (r *Retangulo) SetWidth(width float64) {
	r.width = width
}

func (r *Retangulo) SetHeight(height float64) {
	r.height = height
}

func main() {
	quadrado := &Quadrado{}
	quadrado.SetSize(2.0)

	retangulo := &Retangulo{}
	retangulo.SetWidth(3.0)
	retangulo.SetHeight(3.0)

	fmt.Println(quadrado.GetArea())
	fmt.Println(retangulo.GetArea())
}
