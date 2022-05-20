package bothub

import (
	"math"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (vector Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		X: vector.X + other.X,
		Y: vector.Y + other.Y,
		Z: vector.Z + other.Z,
	}
}

func (vector Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		X: vector.X - other.X,
		Y: vector.Y - other.Y,
		Z: vector.Z - other.Z,
	}
}

func (vector Vector3) Scale(n float64) Vector3 {
	vector.X *= n
	vector.Y *= n
	vector.Z *= n
	return vector
}

func (vector Vector3) Length() float64 {
	return math.Sqrt(vector.Dot(vector))
}

func Squared(n float64) float64 {
	return math.Pow(n, 2)
}

func ThreeDimensionalEuclidianDistance(point1 Vector3, point2 Vector3) float64 {
	return math.Sqrt(Squared(point1.X-point2.X) + Squared(point1.Y-point2.Y) + Squared(point1.Z-point2.Z))
}

func (vector Vector3) Dot(other Vector3) float64 {
	return vector.X*other.X + vector.Y*other.Y + vector.Z*other.Z
}

func (vector Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		X: vector.Y*other.Z - vector.Z*other.Y,
		Y: vector.Z*other.X - vector.X*other.Z,
		Z: vector.X*other.Y - vector.Y*other.X,
	}
}

func (vector Vector3) Normalize() Vector3 {
	vector.Scale(1 / vector.Length())
	return vector
}

func MaxOfTwoVectorsInEachDimension(distance1 Vector3, distance2 Vector3) Vector3 {
	return Vector3{
		X: math.Max(distance1.X, distance2.X),
		Y: math.Max(distance1.Y, distance2.Y),
		Z: math.Max(distance1.Z, distance2.Z),
	}
}
