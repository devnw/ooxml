package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideLayoutIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_SlideLayoutIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_SlideLayoutIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideLayoutIdListEntry should validate: %s", err)
	}
}

func TestCT_SlideLayoutIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideLayoutIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideLayoutIdListEntry()
	xml.Unmarshal(buf, v2)
}
