package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_VolMainConstructor(t *testing.T) {
	v := sml.NewCT_VolMain()
	if v == nil {
		t.Errorf("sml.NewCT_VolMain must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_VolMain should validate: %s", err)
	}
}

func TestCT_VolMainMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_VolMain()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_VolMain()
	xml.Unmarshal(buf, v2)
}
