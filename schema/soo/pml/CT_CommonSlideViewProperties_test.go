package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommonSlideViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_CommonSlideViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_CommonSlideViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommonSlideViewProperties should validate: %s", err)
	}
}

func TestCT_CommonSlideViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommonSlideViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommonSlideViewProperties()
	xml.Unmarshal(buf, v2)
}
