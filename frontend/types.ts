interface Vector3 {
    x: number
    y: number
    z: number
}

interface Point {
	position:              Vector3
	velocity:              Vector3
	bounciness:            number
	pinned:                boolean
	positionReset:         Vector3
}

interface Stick {
    a: Point
    b: Point
    size: Point
}

export {Vector3, Point, Stick}