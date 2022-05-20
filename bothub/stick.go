package bothub

type Stick struct {
	A    *Point
	B    *Point
	Size float64
}

func MakeStick(a *Point, b *Point) Stick {
	vTmp := b.Position.Sub(a.Position)
	return Stick{
		A:    a,
		B:    b,
		Size: vTmp.Length(),
	}
}

func (s *Stick) Update() {
	distance := ThreeDimensionalEuclidianDistance(s.A.Position, s.B.Position)
	dL := (s.Size - distance) / distance / 2
	offset := s.A.Position.Sub(s.B.Position)
	scaled_offset := offset.Scale(dL)
	if !s.A.Pinned {
		s.A.Position = s.A.Position.Add(scaled_offset)
	}
	if !s.B.Pinned {
		s.B.Position = s.B.Position.Sub(scaled_offset)
	}
}
