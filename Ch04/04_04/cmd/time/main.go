package main

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	t := time.Date(2024, 4, 13, 17, 32, 47, 203, time.UTC)
	fmt.Println("t :", t)

	// time.Time -> timestamppb.Timestamp
	pt := timestamppb.New(t)
	fmt.Println("pt:", pt)

	// timestamppb.Timestamp -> time.Time
	t2 := pt.AsTime()
	fmt.Println("t2:", t2)
}
