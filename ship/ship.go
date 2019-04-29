package ship

// Ship contains the data for a space ship
type Ship struct {
	Name           string
	Shield         Defense
	Armor          Defense
	Hull           Defense
	MaxTargetRange int
	Capacitor      Capacitor
	Weapons        []Weapon
	MaxVelocity    int
	CurVelocity    int
	Inertia        float64
	Mass           int
	Radius         int
}

type Space struct {
	Ship     Ship
	Location Location
}

func (s Space) updateLocation() {

}

// Location of the ship on the grid
type Location struct {
	x, y int
}

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
