package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{
		sp: sp,
	}
}

type IceSafetyPlacer struct{}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my ICE safeties...")
}

type SandSafetyPlacer struct{}

func (sp SandSafetyPlacer) placeSafeties() {
	fmt.Println("placing my SAND safeties...")
}

type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("placing NO safeties...")
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc1 := newRockClimber(IceSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc1.climbRock()
	}

	rc2 := newRockClimber(SandSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc2.climbRock()
	}

	rc3 := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 11; i++ {
		rc3.climbRock()
	}
}
