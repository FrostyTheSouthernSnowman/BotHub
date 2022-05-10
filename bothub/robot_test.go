package bothub

import (
	"math"
	"testing"
)

func TestEuclidianDistance(t *testing.T) {
	point1 := Vector3{X: 1, Y: 1, Z: 1}
	point2 := Vector3{X: 2, Y: 3, Z: 2}
	distance := ThreeDimensionalEuclidianDistance(point1, point2)
	if distance != float32(math.Sqrt(6)) {
		t.Fatalf("distance should be one, got %v", distance)
	}
}
