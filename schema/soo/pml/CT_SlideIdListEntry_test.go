package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_SlideIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_SlideIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideIdListEntry should validate: %s", err)
	}
}

func TestCT_SlideIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideIdListEntry()
	xml.Unmarshal(buf, v2)
}
