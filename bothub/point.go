package bothub

type Point struct {
	Position              Vector3
	prevPosition          Vector3
	Velocity              Vector3
	Bounciness            float64
	Friction              float64
	Pinned                bool
	PositionReset         Vector3
	previousPositionReset Vector3
}

// MakePoint takes in a position, bounciness, and friction and returns a point
//
// It is recommended to use bounciness of 0.9 and friction of 0.99
func MakePoint(position Vector3, bounciness float64, friction float64) Point {
	return Point{
		Position:     position,
		prevPosition: position,
		Velocity: Vector3{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Bounciness:            bounciness,
		Friction:              friction,
		Pinned:                false,
		PositionReset:         position,
		previousPositionReset: position,
	}
}

func (p *Point) Reset() {
	p.Position = p.PositionReset
}

func (p *Point) Update(delta float64) {
	if p.Pinned {
		return
	}

	// Retrieve velocity as delta of space
	p.Velocity = p.Position.Sub(p.prevPosition).Scale(1 / delta) // ds/dt = derivative of space with respect to time
	// Apply friction
	p.Velocity = p.Velocity.Scale(p.Friction)

	p.prevPosition = p.Position

	p.Position = p.Position.Add(Vector3{0, 0, -9.8 * delta})

	// Apply velocity
	p.Position = p.Position.Add(p.Velocity.Scale(delta))

	if p.Position.Z < 0 {
		p.Position.Z = 0
		p.Velocity.Z *= -1
		p.Velocity.Z *= p.Bounciness
	}
}
