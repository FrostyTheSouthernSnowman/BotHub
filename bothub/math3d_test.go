package bothub

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScale(t *testing.T) {
	v := Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}
	v = v.Scale(2)
	v2 := Vector3{
		X: 2,
		Y: 4,
		Z: 6,
	}
	if v != v2 {
		t.Fatalf("v: %v, v2: %v", v, v2)
	}
}

func TestAddVectors(t *testing.T) {
	v1 := Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}
	v2 := Vector3{
		X: 2,
		Y: 4,
		Z: 6,
	}
	v3 := Vector3{
		X: 3,
		Y: 6,
		Z: 9,
	}
	sum := v1.Add(v2)

	if sum != v3 {
		t.Fatalf("sum: %v, correct: %v", sum, v3)
	}
}

func TestSubtractVectors(t *testing.T) {
	v1 := Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}
	v2 := Vector3{
		X: 2,
		Y: 4,
		Z: 6,
	}
	v3 := Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}
	diff := v2.Sub(v1)

	if diff != v3 {
		t.Fatalf("sum: %v, correct: %v", diff, v3)
	}
}

func TestSquared(t *testing.T) {
	var a float64 = 5
	var a_squared = Squared(a)
	if a_squared != 25 {
		t.Fatalf("result: %v, correct: %v", a_squared, 25)
	}
}

func TestLength(t *testing.T) {
	v := Vector3{
		X: 1,
		Y: 2,
		Z: 3,
	}

	var length_of_v float64 = v.Length()

	var sqrt_14 float64 = math.Sqrt(14)

	if length_of_v != sqrt_14 {
		t.Fatalf("result: %v, correct: %v", length_of_v, sqrt_14)
	}
}

func TestEuclidianDistance(t *testing.T) {
	point1 := Vector3{X: 1, Y: 1, Z: 1}
	point2 := Vector3{X: 2, Y: 3, Z: 2}
	distance := ThreeDimensionalEuclidianDistance(point1, point2)
	if distance != float64(math.Sqrt(6)) {
		t.Fatalf("distance should be one, got %v", distance)
	}
}

func TestDotProductOfTwoPerpendicularVectors(t *testing.T) {
	result := Vector3{1., 0., 0.}.Dot(Vector3{0., 1., 0.})
	if result != 0 {
		t.Fatalf("result: %v, correct: %v", result, 0)
	}
}

func TestDotProductOfTwoParallelVectors(t *testing.T) {
	result := Vector3{1., 0., 0.}.Dot(Vector3{1., 0., 0.})
	if result != 1 {
		t.Fatalf("result: %v, correct: %v", result, 1)
	}
}

func TestCrossProduct(t *testing.T) {
	result := Vector3{1., 0., 0.}.Cross(Vector3{0., 1., 0.})
	if result != (Vector3{0., 0., 1.}) {
		t.Fatalf("result: %v, correct: %v", result, Vector3{0, 0, 1})
	}
}

func TestNormalize(t *testing.T) {
	assert.Equal(t, Vector3{1., 0., 0.}, Vector3{10., 0., 0.}.Normalize(),
		"returns a unit vector (vector) from the given vector")
	assert.Equal(t, Vector3{0., 1., 0.}, Vector3{0., 10., 0.}.Normalize(),
		"returns a unit vector (vector) from the given vector")
	assert.Equal(t, Vector3{0., 0., 1.}, Vector3{0., 0., 10.}.Normalize(),
		"returns a unit vector (vector) from the given vector")
}

func TestMaxOfTwoVectorsInEachDimension(t *testing.T) {
	assert.Equal(t, Vector3{3, 4, 5}, MaxOfTwoVectorsInEachDimension(Vector3{3, 2, 4}, Vector3{2, 4, 5}), "Maximum value of each dimension of two vectors")
}
