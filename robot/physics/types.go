package physics

type XYZPosition struct {
	X         float32
	Y         float32
	Z         float32
	XRotation float32
	YRotation float32
	ZRotation float32
	Velocity  Vector3
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}
