package schemaLibrary

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_Schema", NewCT_Schema)
	office.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "CT_SchemaLibrary", NewCT_SchemaLibrary)
	office.RegisterConstructor("http://schemas.openxmlformats.org/schemaLibrary/2006/main", "schemaLibrary", NewSchemaLibrary)
}
