package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TimeNodeListConstructor(t *testing.T) {
	v := pml.NewCT_TimeNodeList()
	if v == nil {
		t.Errorf("pml.NewCT_TimeNodeList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TimeNodeList should validate: %s", err)
	}
}

func TestCT_TimeNodeListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TimeNodeList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TimeNodeList()
	xml.Unmarshal(buf, v2)
}
