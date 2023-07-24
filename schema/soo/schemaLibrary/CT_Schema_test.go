package schemaLibrary_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/schemaLibrary"
)

func TestCT_SchemaConstructor(t *testing.T) {
	v := schemaLibrary.NewCT_Schema()
	if v == nil {
		t.Errorf("schemaLibrary.NewCT_Schema must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed schemaLibrary.CT_Schema should validate: %s", err)
	}
}

func TestCT_SchemaMarshalUnmarshal(t *testing.T) {
	v := schemaLibrary.NewCT_Schema()
	buf, _ := xml.Marshal(v)
	v2 := schemaLibrary.NewCT_Schema()
	xml.Unmarshal(buf, v2)
}
