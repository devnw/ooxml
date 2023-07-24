package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ParametersConstructor(t *testing.T) {
	v := sml.NewCT_Parameters()
	if v == nil {
		t.Errorf("sml.NewCT_Parameters must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Parameters should validate: %s", err)
	}
}

func TestCT_ParametersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Parameters()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Parameters()
	xml.Unmarshal(buf, v2)
}
