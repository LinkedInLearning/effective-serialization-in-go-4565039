package elevator

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

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
	var e map[string]any
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}

	switch e["type"] {
	case "door":
		var d DoorEvent
		if err := mapstructure.Decode(e, &d); err != nil {
			return err
		}
		fmt.Println("door event:", d) // TODO: Handle event
	case "button":
		var b ButtonEvent
		if err := mapstructure.Decode(e, &b); err != nil {
			return err
		}
		fmt.Println("button event:", b) // TODO: Handle event
	default:
		return fmt.Errorf("unknown event - %q", e["type"])
	}

	return nil
}
