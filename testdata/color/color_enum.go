// GENERATED BY magnum, DO NOT EDIT

package color

// Colors returns all possible Colors.
func Colors() []Color {
	return []Color{Red(), Orange(), Yellow(), Green(), Blue(), Indigo(), Violet()}
}

// Red returns the "red" Color.
func Red() Color {
	return Color{
		name: "red",
	}
}

// Orange returns the "orange" Color.
func Orange() Color {
	return Color{
		name: "orange",
	}
}

// Yellow returns the "yellow" Color.
func Yellow() Color {
	return Color{
		name: "yellow",
	}
}

// Green returns the "green" Color.
func Green() Color {
	return Color{
		name: "green",
	}
}

// Blue returns the "blue" Color.
func Blue() Color {
	return Color{
		name: "blue",
	}
}

// Indigo returns the "indigo" Color.
func Indigo() Color {
	return Color{
		name: "indigo",
	}
}

// Violet returns the "violet" Color.
func Violet() Color {
	return Color{
		name: "violet",
	}
}

// String returns the Color's name.
func (c Color) String() string {
	return c.name
}