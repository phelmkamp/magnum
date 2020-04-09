package color_test

import (
	"fmt"

	"github.com/phelmkamp/magnum/testdata/color"
)

func Example() {
	fmt.Println("Colors() =", color.Colors())

	c := color.Red()
	fmt.Println("c =", c)
	fmt.Println("(c == Red()) =", c == color.Red())
	fmt.Println("(c == Orange()) =", c == color.Orange())
	// Output:
	// Colors() = [red orange yellow green blue indigo violet]
	// c = red
	// (c == Red()) = true
	// (c == Orange()) = false
}
