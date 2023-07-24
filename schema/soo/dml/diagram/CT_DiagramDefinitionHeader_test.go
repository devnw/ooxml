package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DiagramDefinitionHeaderConstructor(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeader()
	if v == nil {
		t.Errorf("diagram.NewCT_DiagramDefinitionHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DiagramDefinitionHeader should validate: %s", err)
	}
}

func TestCT_DiagramDefinitionHeaderMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DiagramDefinitionHeader()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DiagramDefinitionHeader()
	xml.Unmarshal(buf, v2)
}
