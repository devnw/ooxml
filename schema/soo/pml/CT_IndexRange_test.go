package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_IndexRangeConstructor(t *testing.T) {
	v := pml.NewCT_IndexRange()
	if v == nil {
		t.Errorf("pml.NewCT_IndexRange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_IndexRange should validate: %s", err)
	}
}

func TestCT_IndexRangeMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_IndexRange()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_IndexRange()
	xml.Unmarshal(buf, v2)
}
