package house

import "github.com/phelmkamp/magnum/testdata/color"

type House struct {
	name    string      `enum:"Gryffindor,Hufflepuff,Ravenclaw,Slytherin"`
	color   color.Color `enum:"Red(),Yellow(),Violet(),Green()"`
	founder string      `enum:"Godric Gryffindor,Helga Hufflepuff,Rowena Ravenclaw,Salazar Slytherin"`
}
