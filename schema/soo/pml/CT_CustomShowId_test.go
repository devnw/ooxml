package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CustomShowIdConstructor(t *testing.T) {
	v := pml.NewCT_CustomShowId()
	if v == nil {
		t.Errorf("pml.NewCT_CustomShowId must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CustomShowId should validate: %s", err)
	}
}

func TestCT_CustomShowIdMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CustomShowId()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CustomShowId()
	xml.Unmarshal(buf, v2)
}
