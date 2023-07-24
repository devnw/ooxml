package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideSyncPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_SlideSyncProperties()
	if v == nil {
		t.Errorf("pml.NewCT_SlideSyncProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideSyncProperties should validate: %s", err)
	}
}

func TestCT_SlideSyncPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideSyncProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideSyncProperties()
	xml.Unmarshal(buf, v2)
}
