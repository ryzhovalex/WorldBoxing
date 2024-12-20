package world

import "worldboxing/lib/utils"

type World struct {
	Id utils.Id
	// Viable to keep track on current level of synchronization between events
	// and database state.
	LastProcessedStateEventId utils.Id
}
