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
	X1, Y1 float64
	X2, Y2 float64
}

// CalculateDistance расчитывает расстояние между двумя точками на плоскости
func CalculateDistance(x1, y1, x2, y2 float64) (distance float64, err error) {

	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return 0, errors.New(errNegativeMeaning)
	}

	// возврат расстояния между точками
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)), nil
}
