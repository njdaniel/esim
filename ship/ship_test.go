package ship

import "testing"

func TestVector_Add(t *testing.T) {
	type args struct {
		b Vector
	}
	tests := []struct {
		name string
		a    *Vector
		args args
	}{
			{"positive", &Vector{0,0}, args{Vector{2.0, 3.0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Add(tt.args.b)
		})
	}
}
