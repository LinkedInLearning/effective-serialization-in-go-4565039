package main

import (
	"encoding/json"
	"fmt"
	"serialization/Ch04/04_05/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	msg := pb.RideStart{
		Time:  timestamppb.Now(),
		CarId: "McQueen",
		Location: &pb.Location{
			Lat: 48.8737917,
			Lng: 2.2950275,
		},
		Type:       pb.RideType_SHARED,
		Passengers: 1,
	}

	data, err := json.MarshalIndent(&msg, "", "    ")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}
