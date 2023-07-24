package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IconSetConstructor(t *testing.T) {
	v := sml.NewCT_IconSet()
	if v == nil {
		t.Errorf("sml.NewCT_IconSet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IconSet should validate: %s", err)
	}
}

func TestCT_IconSetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IconSet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IconSet()
	xml.Unmarshal(buf, v2)
}
