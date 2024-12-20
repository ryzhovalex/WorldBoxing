package sevent

import (
	"worldboxing/internal/timeline"
	"worldboxing/lib/utils"
)

type StateEventType = string

const (
	PersonCreated StateEventType = "PersonCreated"
)

type StateEvent struct {
	Id          utils.Id
	Type        string
	Body        string
	Time        utils.Time
	TimelineDay timeline.Day
}

func Create(t StateEventType, bodyMap map[string]any)
