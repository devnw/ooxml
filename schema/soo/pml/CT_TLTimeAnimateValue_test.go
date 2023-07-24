package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TLTimeAnimateValueConstructor(t *testing.T) {
	v := pml.NewCT_TLTimeAnimateValue()
	if v == nil {
		t.Errorf("pml.NewCT_TLTimeAnimateValue must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTimeAnimateValue should validate: %s", err)
	}
}

func TestCT_TLTimeAnimateValueMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTimeAnimateValue()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTimeAnimateValue()
	xml.Unmarshal(buf, v2)
}
