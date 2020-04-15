package house_test

import (
	"encoding/json"
	"fmt"

	"github.com/phelmkamp/magnum/testdata/house"
)

func Example() {
	h := house.Ravenclaw()
	fmt.Println("h =", h)
	fmt.Println("h.Color() =", h.Color())
	fmt.Println("h.Founder() =", h.Founder())

	data, _ := json.Marshal(h)
	fmt.Println("data =", string(data))

	var h2 house.House
	json.Unmarshal(data, &h2)
	fmt.Println("h2 =", h2)
	// Output:
	// h = Ravenclaw
	// h.Color() = violet
	// h.Founder() = Rowena Ravenclaw
	// data = "Ravenclaw"
	// h2 = Ravenclaw
}
