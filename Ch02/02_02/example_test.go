package coord

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

func ExampleCoord_MarshalText() {
	c := Coord{51.523767, -0.1611306}
	if err := json.NewEncoder(os.Stdout).Encode(c); err != nil {
		fmt.Println("error (json):", err)
		return
	}

	if err := xml.NewEncoder(os.Stdout).Encode(c); err != nil {
		fmt.Println("error (xml):", err)
		return
	}

	// Output:
	// "51.523767N 0.161131W"
	// <Coord>51.523767N 0.161131W</Coord>
}
