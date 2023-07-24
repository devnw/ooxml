package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ErrorConstructor(t *testing.T) {
	v := sml.NewCT_Error()
	if v == nil {
		t.Errorf("sml.NewCT_Error must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Error should validate: %s", err)
	}
}

func TestCT_ErrorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Error()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Error()
	xml.Unmarshal(buf, v2)
}
