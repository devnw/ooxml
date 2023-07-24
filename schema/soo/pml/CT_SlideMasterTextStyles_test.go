package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideMasterTextStylesConstructor(t *testing.T) {
	v := pml.NewCT_SlideMasterTextStyles()
	if v == nil {
		t.Errorf("pml.NewCT_SlideMasterTextStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideMasterTextStyles should validate: %s", err)
	}
}

func TestCT_SlideMasterTextStylesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideMasterTextStyles()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideMasterTextStyles()
	xml.Unmarshal(buf, v2)
}
