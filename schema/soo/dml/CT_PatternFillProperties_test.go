package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PatternFillPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_PatternFillProperties()
	if v == nil {
		t.Errorf("dml.NewCT_PatternFillProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PatternFillProperties should validate: %s", err)
	}
}

func TestCT_PatternFillPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PatternFillProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PatternFillProperties()
	xml.Unmarshal(buf, v2)
}
