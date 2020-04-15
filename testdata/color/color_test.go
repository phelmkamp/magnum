package color_test

import (
	"fmt"

	"github.com/phelmkamp/magnum/testdata/color"
)

func Example() {
	fmt.Println("Colors() =", color.Colors())

	c, _ := color.NewColor("red")
	fmt.Println("c =", c)
	fmt.Println("(c == Red()) =", c == color.Red())
	fmt.Println("(c == Orange()) =", c == color.Orange())

	_, err := color.NewColor("blurple")
	fmt.Println("err =", err)
	// Output:
	// Colors() = [red orange yellow green blue indigo violet]
	// c = red
	// (c == Red()) = true
	// (c == Orange()) = false
	// err = unknown name: blurple
}
