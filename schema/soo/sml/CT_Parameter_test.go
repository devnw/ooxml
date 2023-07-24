package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ParameterConstructor(t *testing.T) {
	v := sml.NewCT_Parameter()
	if v == nil {
		t.Errorf("sml.NewCT_Parameter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Parameter should validate: %s", err)
	}
}

func TestCT_ParameterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Parameter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Parameter()
	xml.Unmarshal(buf, v2)
}
