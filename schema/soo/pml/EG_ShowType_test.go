package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestEG_ShowTypeConstructor(t *testing.T) {
	v := pml.NewEG_ShowType()
	if v == nil {
		t.Errorf("pml.NewEG_ShowType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_ShowType should validate: %s", err)
	}
}

func TestEG_ShowTypeMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_ShowType()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_ShowType()
	xml.Unmarshal(buf, v2)
}
