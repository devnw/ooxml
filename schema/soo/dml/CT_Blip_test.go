package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BlipConstructor(t *testing.T) {
	v := dml.NewCT_Blip()
	if v == nil {
		t.Errorf("dml.NewCT_Blip must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Blip should validate: %s", err)
	}
}

func TestCT_BlipMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Blip()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Blip()
	xml.Unmarshal(buf, v2)
}
