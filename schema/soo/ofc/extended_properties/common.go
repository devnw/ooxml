package extended_properties

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_Properties", NewCT_Properties)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorVariant", NewCT_VectorVariant)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_VectorLpstr", NewCT_VectorLpstr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "CT_DigSigBlob", NewCT_DigSigBlob)
	office.RegisterConstructor("http://schemas.openxmlformats.org/officeDocument/2006/extended-properties", "Properties", NewProperties)
}
