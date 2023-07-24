package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SchemaConstructor(t *testing.T) {
	v := sml.NewCT_Schema()
	if v == nil {
		t.Errorf("sml.NewCT_Schema must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Schema should validate: %s", err)
	}
}

func TestCT_SchemaMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Schema()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Schema()
	xml.Unmarshal(buf, v2)
}
