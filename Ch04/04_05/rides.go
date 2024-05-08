package rides

import (
	"fmt"
	"time"
)

type RideType byte

const (
	RegularType RideType = iota + 1
	SharedType
)

func (t RideType) String() string {
	switch t {
	case RegularType:
		return "regular"
	case SharedType:
		return "shared"
	}

	return fmt.Sprintf("RideType(%d)", t)
}

type Location struct {
	Lat float64
	Lng float64
}

type RideStart struct {
	Time       time.Time
	CarID      string
	Location   Location
	Type       RideType
	Passengers int
}
