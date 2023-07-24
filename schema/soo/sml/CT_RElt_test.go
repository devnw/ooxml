package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_REltConstructor(t *testing.T) {
	v := sml.NewCT_RElt()
	if v == nil {
		t.Errorf("sml.NewCT_RElt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RElt should validate: %s", err)
	}
}

func TestCT_REltMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RElt()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RElt()
	xml.Unmarshal(buf, v2)
}
