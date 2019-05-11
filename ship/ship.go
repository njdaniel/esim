package ship

import "math"

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
	Location       Vector
	Heading        Degrees
}

// FullSpeed function to increase Current velocity to Max
func (s *Ship) FullSpeed() {
	s.CurVelocity = s.MaxVelocity
}

func (s *Ship) Stop() {
	s.CurVelocity = 0.0
}

func (s *Ship) Turn(d Degrees) {
	s.Heading += d
	s.Velocity.X = s.CurVelocity * math.Cos(float64(d)*RadToDeg)
	s.Velocity.Y = s.CurVelocity * math.Sin(float64(d)*RadToDeg)
}

// UpateLocation the current velocity updates the current location
func (s *Ship) UpdateLocation() {

}

// Space is a local area containing ships and possibly other objects
type Space []Ship

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
