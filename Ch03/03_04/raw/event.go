package elevator

import (
	"encoding/json"
	"fmt"
)

// Event is the main event type.
type Event struct {
	Type    string
	Payload json.RawMessage
}

// DoorEvent is a door open/close event.
type DoorEvent struct {
	Action string
	Floor  int
}

// ButtonEvent is a button click event.
type ButtonEvent struct {
	Button int
}

// Handle event handles an event.
func HandleEvent(data []byte) error {
	var e Event
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	switch e.Type {
	case "door":
		var d DoorEvent
		if err := json.Unmarshal(e.Payload, &d); err != nil {
			return err
		}
		fmt.Println("door event:", d) // TODO: Handle event
	case "button":
		var b ButtonEvent
		if err := json.Unmarshal(e.Payload, &b); err != nil {
			return err
		}
		fmt.Println("button event:", b) // TODO: Handle event
	default:
		return fmt.Errorf("unknown event - %q", e.Type)
	}

	return nil
}
