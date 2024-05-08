package main

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"serialization/Ch04/04_03/pb"
)

func main() {
	// Marshal
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
	fmt.Println("msg :", &msg)

	data, err := proto.Marshal(&msg)
	if err != nil {
		fmt.Println("ERROR: marshal:", err)
		return
	}
	fmt.Printf("proto size: %5d\n", len(data))

	// Compare size to JSON
	jdata, err := json.Marshal(&msg)
	if err != nil {
		fmt.Println("ERROR: json marshal:", err)
		return
	}
	fmt.Printf(" json size: %5d\n", len(jdata))

	// Unmarshal
	var msg2 pb.RideStart
	if err := proto.Unmarshal(data, &msg2); err != nil {
		fmt.Println("ERROR: unmarshal:", err)
		return
	}
	fmt.Println("msg2:", &msg2)
}
