package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DiagramDefinitionHeaderLstConstructor(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeaderLst()
	if v == nil {
		t.Errorf("diagram.NewCT_DiagramDefinitionHeaderLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DiagramDefinitionHeaderLst should validate: %s", err)
	}
}

func TestCT_DiagramDefinitionHeaderLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeaderLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DiagramDefinitionHeaderLst()
	xml.Unmarshal(buf, v2)
}
