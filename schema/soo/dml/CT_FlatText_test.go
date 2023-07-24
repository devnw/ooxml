package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FlatTextConstructor(t *testing.T) {
	v := dml.NewCT_FlatText()
	if v == nil {
		t.Errorf("dml.NewCT_FlatText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FlatText should validate: %s", err)
	}
}

func TestCT_FlatTextMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FlatText()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FlatText()
	xml.Unmarshal(buf, v2)
}
