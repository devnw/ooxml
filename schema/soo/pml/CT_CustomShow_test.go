package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CustomShowConstructor(t *testing.T) {
	v := pml.NewCT_CustomShow()
	if v == nil {
		t.Errorf("pml.NewCT_CustomShow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CustomShow should validate: %s", err)
	}
}

func TestCT_CustomShowMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CustomShow()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CustomShow()
	xml.Unmarshal(buf, v2)
}
