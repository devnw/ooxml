package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ElemPropSetConstructor(t *testing.T) {
	v := diagram.NewCT_ElemPropSet()
	if v == nil {
		t.Errorf("diagram.NewCT_ElemPropSet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ElemPropSet should validate: %s", err)
	}
}

func TestCT_ElemPropSetMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ElemPropSet()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ElemPropSet()
	xml.Unmarshal(buf, v2)
}
