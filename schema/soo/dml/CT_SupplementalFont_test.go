package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_SupplementalFontConstructor(t *testing.T) {
	v := dml.NewCT_SupplementalFont()
	if v == nil {
		t.Errorf("dml.NewCT_SupplementalFont must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_SupplementalFont should validate: %s", err)
	}
}

func TestCT_SupplementalFontMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_SupplementalFont()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_SupplementalFont()
	xml.Unmarshal(buf, v2)
}
