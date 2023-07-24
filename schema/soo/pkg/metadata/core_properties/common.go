package core_properties

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_CoreProperties", NewCT_CoreProperties)
	office.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keywords", NewCT_Keywords)
	office.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "CT_Keyword", NewCT_Keyword)
	office.RegisterConstructor("http://schemas.openxmlformats.org/package/2006/metadata/core-properties", "coreProperties", NewCoreProperties)
}
