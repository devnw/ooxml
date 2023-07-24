package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_XfConstructor(t *testing.T) {
	v := sml.NewCT_Xf()
	if v == nil {
		t.Errorf("sml.NewCT_Xf must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Xf should validate: %s", err)
	}
}

func TestCT_XfMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Xf()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Xf()
	xml.Unmarshal(buf, v2)
}
