package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_BackgroundConstructor(t *testing.T) {
	v := pml.NewCT_Background()
	if v == nil {
		t.Errorf("pml.NewCT_Background must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Background should validate: %s", err)
	}
}

func TestCT_BackgroundMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Background()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Background()
	xml.Unmarshal(buf, v2)
}
