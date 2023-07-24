package custom_properties

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Properties", NewCT_Properties)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "CT_Property", NewCT_Property)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/custom-properties", "Properties", NewProperties)
}
