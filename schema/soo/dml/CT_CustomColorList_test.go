package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_CustomColorListConstructor(t *testing.T) {
	v := dml.NewCT_CustomColorList()
	if v == nil {
		t.Errorf("dml.NewCT_CustomColorList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_CustomColorList should validate: %s", err)
	}
}

func TestCT_CustomColorListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_CustomColorList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_CustomColorList()
	xml.Unmarshal(buf, v2)
}
