package physics

func CalculatePhysics(simulation []XYZPosition) ([]XYZPosition, error) {
	var newSimulation []XYZPosition
	for _, value := range simulation {
		if CheckIfNotTouchingFloor(value) {
			value, _ = AddGravity(value)
		}
		value.Z += value.Velocity.Z
		if value.Z < 0.5 {
			value.Z = 0.5
		}
		value.Y += value.Velocity.Y
		value.X += value.Velocity.X
		newSimulation = append(newSimulation, value)
	}
	return newSimulation, nil
}
