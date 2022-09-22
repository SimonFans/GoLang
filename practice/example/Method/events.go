package Method

import (
	"fmt"
	"time"
)

type Event struct {
	ID   string
	Time time.Time
}

type DoorEvent struct {
	Event
	Action string // open, close
}

type TemperatureEvent struct {
	Event
	Value float64
}

func NewDoorEvent(id string, time time.Time, action string) (*DoorEvent, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id")
	}

	evt := DoorEvent{
		Event:  Event{id, time},
		Action: action,
	}
	fmt.Printf("%p\n", &evt)
	return &evt, nil
}
