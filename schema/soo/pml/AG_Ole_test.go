package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestAG_OleConstructor(t *testing.T) {
	v := pml.NewAG_Ole()
	if v == nil {
		t.Errorf("pml.NewAG_Ole must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.AG_Ole should validate: %s", err)
	}
}

func TestAG_OleMarshalUnmarshal(t *testing.T) {
	v := pml.NewAG_Ole()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewAG_Ole()
	xml.Unmarshal(buf, v2)
}
