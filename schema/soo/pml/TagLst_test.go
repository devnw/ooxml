package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestTagLstConstructor(t *testing.T) {
	v := pml.NewTagLst()
	if v == nil {
		t.Errorf("pml.NewTagLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.TagLst should validate: %s", err)
	}
}

func TestTagLstMarshalUnmarshal(t *testing.T) {
	v := pml.NewTagLst()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewTagLst()
	xml.Unmarshal(buf, v2)
}
