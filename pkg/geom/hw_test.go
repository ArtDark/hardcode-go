package geom

import (
	"errors"
	"testing"
)

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         Plain
		wantDistance float64
		wantErr      error
	}{

		{
			name: "#1",
			geom: Plain{
				A: Point{
					x: 1,
					y: 1,
				},
				B: Point{
					x: 4,
					y: 5,
				},
			},
			wantDistance: 5,
			wantErr:      nil,
		},
		{
			name: "#2",
			geom: Plain{
				A: Point{
					x: -1,
					y: 1,
				},
				B: Point{
					x: 4,
					y: 5,
				},
			},
			wantDistance: 0,
			wantErr:      errors.New(errNegativeMeaning),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance, err := tt.geom.CalculateDistance(); gotDistance != tt.wantDistance && err != tt.wantErr {
				t.Errorf("Error = %v, want %v", err, nil)
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)

			}

		})
	}
}
