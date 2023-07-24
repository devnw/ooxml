package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotFieldsConstructor(t *testing.T) {
	v := sml.NewCT_PivotFields()
	if v == nil {
		t.Errorf("sml.NewCT_PivotFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotFields should validate: %s", err)
	}
}

func TestCT_PivotFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotFields()
	xml.Unmarshal(buf, v2)
}
