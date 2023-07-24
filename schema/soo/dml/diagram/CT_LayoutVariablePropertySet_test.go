package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_LayoutVariablePropertySetConstructor(t *testing.T) {
	v := diagram.NewCT_LayoutVariablePropertySet()
	if v == nil {
		t.Errorf("diagram.NewCT_LayoutVariablePropertySet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_LayoutVariablePropertySet should validate: %s", err)
	}
}

func TestCT_LayoutVariablePropertySetMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_LayoutVariablePropertySet()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_LayoutVariablePropertySet()
	xml.Unmarshal(buf, v2)
}
