package Method

import "log"

type ClickEvent struct {
	// ...
}

type HoverEvent struct {
	// ...
}

var EventCounts = make(map[string]int) // type -> count

func RecordEvent(evt interface{}) {
	switch evt.(type) {
	case *ClickEvent:
		EventCounts["click"]++
	case *HoverEvent:
		EventCounts["hover"]++
	default:
		log.Printf("warning: unknown event: %#v of type %T\n", evt, evt)
	}
}
