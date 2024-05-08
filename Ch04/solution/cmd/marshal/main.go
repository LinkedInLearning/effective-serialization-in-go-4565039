package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"serialization/Ch04/solution/pb"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// end ride message
	// id: "aa08deb"
	// time: 2024-04-18T07:52:31Z
	// distance: 1.7
	// location: 48.8737820, 2.2950183

	// Save to a file called "end.pb", then load from it.

	t := time.Date(2024, time.April, 18, 7, 52, 31, 0, time.UTC)
	msg := pb.RideEnd{
		Id:       "aa08deb",
		Time:     timestamppb.New(t),
		Distance: 1.7,
		Location: &pb.Location{
			Lat: 48.8737820,
			Lng: 2.2950183,
		},
	}
	fmt.Println("msg :", &msg)

	data, err := proto.Marshal(&msg)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	const fileName = "end.pb"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	if _, err := io.Copy(file, bytes.NewReader(data)); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer file.Close()

	data, err = io.ReadAll(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var msg2 pb.RideEnd
	if err := proto.Unmarshal(data, &msg2); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("msg2:", &msg2)
}
