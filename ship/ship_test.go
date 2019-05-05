package ship

import (
	"reflect"
	"testing"
)

func TestVector_Add(t *testing.T) {
	type fields struct {
		X float64
		Y float64
	}
	type args struct {
		b Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Vector
	}{
		{"add pos pos to 00", fields{0, 0}, args{Vector{2.0, 3.0}}, Vector{2.0, 3.0}},
		// {"add neg neg to 00"},
		// {"add pos pos to pos post"},
		// {"add pos pos to neg neg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := a.Add(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
