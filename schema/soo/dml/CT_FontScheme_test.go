package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FontSchemeConstructor(t *testing.T) {
	v := dml.NewCT_FontScheme()
	if v == nil {
		t.Errorf("dml.NewCT_FontScheme must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FontScheme should validate: %s", err)
	}
}

func TestCT_FontSchemeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FontScheme()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FontScheme()
	xml.Unmarshal(buf, v2)
}
