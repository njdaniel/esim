package models

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

func TestShip_FullSpeed(t *testing.T) {
	tests := []struct {
		name string
		s    *Ship
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.FullSpeed()
		})
	}
}

func TestShip_Stop(t *testing.T) {
	tests := []struct {
		name string
		s    *Ship
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Stop()
		})
	}
}

func TestShip_Turn(t *testing.T) {
	type args struct {
		d Degrees
	}
	tests := []struct {
		name string
		s    *Ship
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Turn(tt.args.d)
		})
	}
}

func TestShip_UpdateLocation(t *testing.T) {
	tests := []struct {
		name string
		s    *Ship
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.UpdateLocation()
		})
	}
}
