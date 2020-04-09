package house_test

import (
	"fmt"

	"github.com/phelmkamp/magnum/testdata/house"
)

func Example() {
	h := house.Ravenclaw()
	fmt.Println("h =", h)
	fmt.Println("h.Color() =", h.Color())
	fmt.Println("h.Founder() =", h.Founder())
	// Output:
	// h = Ravenclaw
	// h.Color() = violet
	// h.Founder() = Rowena Ravenclaw
}
