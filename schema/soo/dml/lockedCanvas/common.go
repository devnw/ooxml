package lockedCanvas

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/lockedCanvas", "lockedCanvas", NewLockedCanvas)
}
