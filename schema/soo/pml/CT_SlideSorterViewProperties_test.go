package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideSorterViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_SlideSorterViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_SlideSorterViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideSorterViewProperties should validate: %s", err)
	}
}

func TestCT_SlideSorterViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideSorterViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideSorterViewProperties()
	xml.Unmarshal(buf, v2)
}
