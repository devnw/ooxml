package picture

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_PictureNonVisual", NewCT_PictureNonVisual)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "CT_Picture", NewCT_Picture)
	office.RegisterConstructor("http://schemas.openxmlformats.org/drawingml/2006/picture", "pic", NewPic)
}
