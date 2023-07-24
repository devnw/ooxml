package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OutlineViewSlideEntryConstructor(t *testing.T) {
	v := pml.NewCT_OutlineViewSlideEntry()
	if v == nil {
		t.Errorf("pml.NewCT_OutlineViewSlideEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OutlineViewSlideEntry should validate: %s", err)
	}
}

func TestCT_OutlineViewSlideEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OutlineViewSlideEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OutlineViewSlideEntry()
	xml.Unmarshal(buf, v2)
}
