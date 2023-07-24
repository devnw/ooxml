package schemaLibrary_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/schemaLibrary"
)

func TestCT_SchemaLibraryConstructor(t *testing.T) {
	v := schemaLibrary.NewCT_SchemaLibrary()
	if v == nil {
		t.Errorf("schemaLibrary.NewCT_SchemaLibrary must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed schemaLibrary.CT_SchemaLibrary should validate: %s", err)
	}
}

func TestCT_SchemaLibraryMarshalUnmarshal(t *testing.T) {
	v := schemaLibrary.NewCT_SchemaLibrary()
	buf, _ := xml.Marshal(v)
	v2 := schemaLibrary.NewCT_SchemaLibrary()
	xml.Unmarshal(buf, v2)
}
