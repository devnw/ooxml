package schemaLibrary_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/schemaLibrary"
)

func TestSchemaLibraryConstructor(t *testing.T) {
	v := schemaLibrary.NewSchemaLibrary()
	if v == nil {
		t.Errorf("schemaLibrary.NewSchemaLibrary must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed schemaLibrary.SchemaLibrary should validate: %s", err)
	}
}

func TestSchemaLibraryMarshalUnmarshal(t *testing.T) {
	v := schemaLibrary.NewSchemaLibrary()
	buf, _ := xml.Marshal(v)
	v2 := schemaLibrary.NewSchemaLibrary()
	xml.Unmarshal(buf, v2)
}
