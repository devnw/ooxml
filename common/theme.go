package common

import "go.devnw.com/ooxml/schema/soo/dml"

// Theme is a drawingml theme.
type Theme struct {
	x *dml.Theme
}

// NewTheme constructs a new theme.
func NewTheme() Theme {
	return Theme{dml.NewTheme()}
}

// X returns the inner wrapped XML type.
func (t Theme) X() *dml.Theme {
	return t.x
}
