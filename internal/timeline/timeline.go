package timeline

import (
	"worldboxing/lib/database"
	"worldboxing/lib/utils"
)

type Day = int
type Timeline struct {
	Id         utils.Id
	CurrentDay Day
}

var cachedCurrentDay *Day = nil

func CurrentDay() (Day, *utils.Error) {
	if cachedCurrentDay == nil {
		var timeline Timeline
		be := database.T.Select(&timeline, "SELECT * FROM Timeline LIMIT 1")
		if be != nil {
			panic("Unstarted world.")
		}
		cachedCurrentDay = &timeline.CurrentDay
	}
	return *cachedCurrentDay, nil
}
