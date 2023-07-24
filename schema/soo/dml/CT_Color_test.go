package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorConstructor(t *testing.T) {
	v := dml.NewCT_Color()
	if v == nil {
		t.Errorf("dml.NewCT_Color must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Color should validate: %s", err)
	}
}

func TestCT_ColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Color()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Color()
	xml.Unmarshal(buf, v2)
}
