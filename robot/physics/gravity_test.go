package physics

import (
	"reflect"
	"testing"
)

func TestAddGravity(t *testing.T) {
	type args struct {
		object XYZPosition
	}
	tests := []struct {
		name    string
		args    args
		want    XYZPosition
		wantErr bool
	}{
		{
			name: "The velocity should be accelarated by 9.8m/s divided by the number of calculations run per second as defined in the constants.go file",
			args: args{object: XYZPosition{X: 0, Y: 0, Z: 0, Velocity: Vector3{Z: 0}}},
			want: XYZPosition{
				X:         0,
				Y:         0,
				Z:         0,
				XRotation: 0,
				YRotation: 0,
				ZRotation: 0,
				Velocity: Vector3{
					X: 0,
					Y: 0,
					Z: -9.8 / CalculationsPerSecond,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddGravity(tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddGravity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddGravity() = %v, want %v", got, tt.want)
			}
		})
	}
}
