package main

import (
	"fmt"
	"serialization/Ch04/04_05/pb"

	"google.golang.org/protobuf/encoding/protojson"
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
		Type:       pb.RideType_REGULAR,
		Passengers: 1,
	}

	// data, err := json.MarshalIndent(&msg, "", "    ")
	data, err := protojson.Marshal(&msg)

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(string(data))
}
