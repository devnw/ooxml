package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ReviewedConstructor(t *testing.T) {
	v := sml.NewCT_Reviewed()
	if v == nil {
		t.Errorf("sml.NewCT_Reviewed must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Reviewed should validate: %s", err)
	}
}

func TestCT_ReviewedMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Reviewed()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Reviewed()
	xml.Unmarshal(buf, v2)
}
