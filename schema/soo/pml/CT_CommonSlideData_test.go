package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommonSlideDataConstructor(t *testing.T) {
	v := pml.NewCT_CommonSlideData()
	if v == nil {
		t.Errorf("pml.NewCT_CommonSlideData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommonSlideData should validate: %s", err)
	}
}

func TestCT_CommonSlideDataMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommonSlideData()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommonSlideData()
	xml.Unmarshal(buf, v2)
}
