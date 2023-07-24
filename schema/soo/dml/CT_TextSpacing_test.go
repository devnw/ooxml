package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextSpacingConstructor(t *testing.T) {
	v := dml.NewCT_TextSpacing()
	if v == nil {
		t.Errorf("dml.NewCT_TextSpacing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextSpacing should validate: %s", err)
	}
}

func TestCT_TextSpacingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextSpacing()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextSpacing()
	xml.Unmarshal(buf, v2)
}
