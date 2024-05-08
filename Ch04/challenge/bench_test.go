package rides

import (
	"encoding/json"
	"serialization/Ch04/04_03/pb"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var msg = pb.RideStart{
	Time:  timestamppb.Now(),
	CarId: "McQueen",
	Location: &pb.Location{
		Lat: 48.8737917,
		Lng: 2.2950275,
	},
	Type:       pb.RideType_SHARED,
	Passengers: 1,
}

func BenchmarkProto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(&msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}
