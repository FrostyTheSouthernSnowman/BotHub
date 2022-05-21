package bothub

import "math"

func MakeCube(side float64, center Vector3) ([]*Stick, []*Point) {
	var points []*Point
	var sticks []*Stick

	a := MakePoint(center.Sub(Vector3{side / 2, side / 2, side / 2}), 0.9, 0.99)
	b := MakePoint(a.Position.Add(Vector3{side, 0, 0}), 0.9, 0.99)
	c := MakePoint(a.Position.Add(Vector3{side, side, 0}), 0.9, 0.99)
	d := MakePoint(a.Position.Add(Vector3{0, side, 0}), 0.9, 0.99)

	e := MakePoint(a.Position.Add(Vector3{0, 0, side}), 0.9, 0.99)
	f := MakePoint(a.Position.Add(Vector3{side, 0, side}), 0.9, 0.99)
	g := MakePoint(a.Position.Add(Vector3{side, side, side}), 0.9, 0.99)
	h := MakePoint(a.Position.Add(Vector3{0, side, side}), 0.9, 0.99)

	points = append(points, &a)
	points = append(points, &b)
	points = append(points, &c)
	points = append(points, &d)

	points = append(points, &e)
	points = append(points, &f)
	points = append(points, &g)
	points = append(points, &h)

	stick01 := MakeStick(&a, &b)
	stick02 := MakeStick(&b, &c)
	stick03 := MakeStick(&c, &d)
	stick04 := MakeStick(&d, &a)

	stick05 := MakeStick(&e, &f)
	stick06 := MakeStick(&f, &g)
	stick07 := MakeStick(&g, &h)
	stick08 := MakeStick(&h, &e)

	stick09 := MakeStick(&a, &e)
	stick10 := MakeStick(&b, &f)
	stick11 := MakeStick(&c, &g)
	stick12 := MakeStick(&d, &h)

	stick13 := MakeStick(&a, &g)
	stick14 := MakeStick(&b, &h)
	stick15 := MakeStick(&c, &e)
	stick16 := MakeStick(&d, &f)

	stick17 := MakeStick(&a, &c)
	stick18 := MakeStick(&b, &g)
	stick19 := MakeStick(&f, &h)
	stick20 := MakeStick(&d, &e)

	stick21 := MakeStick(&d, &g)
	stick22 := MakeStick(&a, &f)

	sticks = append(sticks, &stick01)
	sticks = append(sticks, &stick02)
	sticks = append(sticks, &stick03)
	sticks = append(sticks, &stick04)

	sticks = append(sticks, &stick05)
	sticks = append(sticks, &stick06)
	sticks = append(sticks, &stick07)
	sticks = append(sticks, &stick08)

	sticks = append(sticks, &stick09)
	sticks = append(sticks, &stick10)
	sticks = append(sticks, &stick11)
	sticks = append(sticks, &stick12)

	sticks = append(sticks, &stick13)
	sticks = append(sticks, &stick14)
	sticks = append(sticks, &stick15)
	sticks = append(sticks, &stick16)

	sticks = append(sticks, &stick17)
	sticks = append(sticks, &stick18)
	sticks = append(sticks, &stick19)
	sticks = append(sticks, &stick20)

	sticks = append(sticks, &stick21)
	sticks = append(sticks, &stick22)

	return sticks, points
}

func MakeSphere(radius float64, center Vector3) ([]*Stick, []*Point) {
	stride := (2 * math.Pi) / 10 * radius

	var points []*Point
	var sticks []*Stick

	// 20 * radius * 0.5 = 10 * radius
	for i := 0.0; i < 10*radius; i++ {
		theta := i * stride

		point1 := MakePoint(Vector3{center.X + math.Cos(theta)*radius, 0, center.Z + math.Sin(theta)*radius}, 0.9, 0.99)
		point2 := MakePoint(Vector3{center.X - math.Cos(theta)*radius, 0, center.Z - math.Sin(theta)*radius}, 0.9, 0.99)

		points = append(points, &point1)
		points = append(points, &point2)

		stick := MakeStick(&point1, &point2)
		sticks = append(sticks, &stick)
	}

	for i := 0; i < len(points); i++ {
		outer_stick := MakeStick(points[i], points[(i+2)%len(points)])
		sticks = append(sticks, &outer_stick)
	}

	return sticks, points
}
