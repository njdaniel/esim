package models

import (
	"fmt"
	"math"
	"sync"
)

const (
	RadToDeg = 180 / math.Pi
	DegToRad = math.Pi / 180
)

// Ship contains the data for a space ship
type Ship struct {
	Name           string
	Shield         Defense
	Armor          Defense
	Hull           Defense
	MaxTargetRange int
	Capacitor      Capacitor
	Weapons        []Weapon
	MaxVelocity    float64
	CurVelocity    float64
	Velocity       Vector
	Inertia        float64
	Mass           int
	Radius         int
	Location       Vector // Probably going to move this to the space object
	Heading        Degrees
	Mutex          sync.Mutex
}

func (s *Ship) Print() {
	fmt.Printf("Ship Name: %s \n"+
		"MaxVelocity: %v \n"+
		"CurVelocity: %v \n"+
		"Velocity: %v \n"+
		"Location: %v \n", s.Name, s.MaxVelocity, s.CurVelocity, s.Velocity, s.Location)
}

// FullSpeed function to increase Current velocity to Max
func (s *Ship) FullSpeed() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.CurVelocity = s.MaxVelocity
}

func (s *Ship) Stop() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.CurVelocity = 0.0
	s.Velocity.X = 0.0
	s.Velocity.Y = 0.0
}

func (s *Ship) Turn(d Degrees) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Heading += d
	s.Velocity.X = s.CurVelocity * math.Cos(float64(s.Heading)*RadToDeg)
	s.Velocity.Y = s.CurVelocity * math.Sin(float64(s.Heading)*RadToDeg)
}

// UpdateLocation the current velocity updates the current location
func (s *Ship) UpdateLocation() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Velocity.X = s.CurVelocity * math.Cos(float64(s.Heading)*RadToDeg)
	s.Velocity.Y = s.CurVelocity * math.Sin(float64(s.Heading)*RadToDeg)
	s.Location.Add(s.Velocity)
}

// Space is a local area containing ships and possibly other objects =================================
type Space map[string]*Ship

// Vector struct holding the x y values of a 2d vector
type Vector struct {
	X, Y float64
}

// Add Two vectors
// TODO: returning Vector for testing.. should i return vector?
// pointer receiver 'a' is already being modified
func (a *Vector) Add(b Vector) Vector {
	a.X += b.X
	a.Y += b.Y
	return *a
}

// Degrees is the direction the ship is pointing towards on the Space grid (0-360)
type Degrees float64

// Defense contains the hitpoints and the resistance
type Defense struct {
	MaxHP    int
	CurHP    int
	EMRes    float64
	ExpRes   float64
	KinRes   float64
	TherRes  float64
	Recharge int
}

// Capacitor is the power source to run the modules such as weapons and reppers. Has the max amount and the recharge rate.
type Capacitor struct {
	Capacity int
	Percent  float64
	Recharge int
}

// Weapon is a module that deals damage in the form of EM, Explosive, Kinetic, and Thermal.
type Weapon struct {
	EMDmg    int
	ExpDmg   int
	KinDmg   int
	TherDmg  int
	Rof      float64
	Optimal  int
	Falloff  int
	Tracking float64
}

// ShieldBooster activated module to recover shield up to max.
type ShieldBooster struct {
	ShieldBonus    int
	ActivationTime float64
	ActivationCost int
}
