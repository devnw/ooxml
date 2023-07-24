package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideMasterConstructor(t *testing.T) {
	v := pml.NewCT_SlideMaster()
	if v == nil {
		t.Errorf("pml.NewCT_SlideMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideMaster should validate: %s", err)
	}
}

func TestCT_SlideMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideMaster()
	xml.Unmarshal(buf, v2)
}
