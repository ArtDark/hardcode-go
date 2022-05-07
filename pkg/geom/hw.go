// Package geom содержит формулы из аналитаической геометрии
package geom

import (
	"errors"
	"math"
)

// errNegativeMeaning По условиям задачи, координаты не могут быть меньше 0.
const errNegativeMeaning = "координаты не могут быть меньше нуля"

// Plain координатная плоскость
type Plain struct {
	A Point
	B Point
}

// Point координаты точки на плоскости
type Point struct {
	x float64
	y float64
}

// NewPlain создает координатную плоскость с двумя точками
func NewPlain(a Point, b Point) *Plain {
	return &Plain{A: a, B: b}
}

// NewPoint создает точку на координатной плоскости
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// CalculateDistance расчитывает расстояние между двумя точками на плоскости
func (p *Plain) CalculateDistance() (distance float64, err error) {

	if p.A.x < 0 || p.B.x < 0 || p.A.y < 0 || p.B.y < 0 {

		return 0, errors.New(errNegativeMeaning)
	} else {
		distance = math.Sqrt(math.Pow(p.B.x-p.A.x, 2) + math.Pow(p.B.y-p.A.y, 2))
	}

	// возврат расстояния между точками
	return distance, nil
}
