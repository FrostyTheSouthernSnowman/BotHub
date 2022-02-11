package physics

func AddGravity(object XYZPosition) (XYZPosition, error) {
	//* Simulated for earth's at 9.8m/s/s of acceleration due to gravity
	object.Velocity.Z -= 9.8 / CalculationsPerSecond
	return object, nil
}
