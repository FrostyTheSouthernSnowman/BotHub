package physics

import "testing"

func TestCheckIfNotTouchingFloor(t *testing.T) {
	type args struct {
		object XYZPosition
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "tests for box under the ground, should return false",
			args: args{object: XYZPosition{X: 0, Y: 0, Z: -5}},
			want: false,
		},
		{
			name: "tests for box above the ground, should return true",
			args: args{object: XYZPosition{X: 0, Y: 0, Z: 5}},
			want: true,
		},
		{
			name: "tests for box touching the ground, should return false",
			args: args{object: XYZPosition{X: 0, Y: 0, Z: 0.5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfNotTouchingFloor(tt.args.object); got != tt.want {
				t.Errorf("CheckIfNotTouchingFloor() = %v, want %v", got, tt.want)
			}
		})
	}
}
