package email

import "fmt"

func ExampleEmail() {
	e := Email{
		Name:    "Dipper Pines",
		Address: "dipper@pines.familiy",
	}
	fmt.Println(e)

	// Output:
	// Dipper Pines <dipper@pines.familiy>
}
