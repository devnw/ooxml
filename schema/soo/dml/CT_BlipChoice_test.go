package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BlipChoiceConstructor(t *testing.T) {
	v := dml.NewCT_BlipChoice()
	if v == nil {
		t.Errorf("dml.NewCT_BlipChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BlipChoice should validate: %s", err)
	}
}

func TestCT_BlipChoiceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BlipChoice()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BlipChoice()
	xml.Unmarshal(buf, v2)
}
