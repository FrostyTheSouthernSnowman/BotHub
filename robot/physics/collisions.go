package physics

func CheckIfNotTouchingFloor(object XYZPosition) bool {
	if object.Z > 0.5 {
		return true
	} else {
		return false
	}
}
