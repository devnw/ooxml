package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SmartTagsConstructor(t *testing.T) {
	v := pml.NewCT_SmartTags()
	if v == nil {
		t.Errorf("pml.NewCT_SmartTags must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SmartTags should validate: %s", err)
	}
}

func TestCT_SmartTagsMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SmartTags()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SmartTags()
	xml.Unmarshal(buf, v2)
}
