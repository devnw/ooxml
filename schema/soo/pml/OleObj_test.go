package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestOleObjConstructor(t *testing.T) {
	v := pml.NewOleObj()
	if v == nil {
		t.Errorf("pml.NewOleObj must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.OleObj should validate: %s", err)
	}
}

func TestOleObjMarshalUnmarshal(t *testing.T) {
	v := pml.NewOleObj()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewOleObj()
	xml.Unmarshal(buf, v2)
}
