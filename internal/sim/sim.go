package sim

import (
	"worldboxing/internal/world"
	"worldboxing/lib/utils"
)

// To populate world for simulation we need to:
// * create countries
// * create cities
// * create companies
func Start() *utils.Error {
	worldObject := world.GetWorld()
	if worldObject.SimulationStarted {
		return utils.DefaultError("Simulation already started.")
	}
	world.SetSimulationStarted()
}
