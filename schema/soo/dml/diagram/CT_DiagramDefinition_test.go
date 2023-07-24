package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DiagramDefinitionConstructor(t *testing.T) {
	v := diagram.NewCT_DiagramDefinition()
	if v == nil {
		t.Errorf("diagram.NewCT_DiagramDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DiagramDefinition should validate: %s", err)
	}
}

func TestCT_DiagramDefinitionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DiagramDefinition()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DiagramDefinition()
	xml.Unmarshal(buf, v2)
}
