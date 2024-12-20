package world

import (
	"worldboxing/lib/database"
	"worldboxing/lib/utils"
)

type World struct {
	Id utils.Id
	// Viable to keep track on current level of synchronization between events
	// and database state.
	LastProcessedStateEventId utils.Id
	SimulationStarted         bool
}

func GetWorld() *World {
	var world World
	be := database.T.Select(&world, "SELECT * FROM World LIMIT 1")
	if be != nil {
		panic("Unstarted world.")
	}
	return &world
}

func SetSimulationStarted() {
	database.T.MustExec("UPDATE World SET SimulationStarted = 1")
}
