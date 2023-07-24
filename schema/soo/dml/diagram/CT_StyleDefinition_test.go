package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_StyleDefinitionConstructor(t *testing.T) {
	v := diagram.NewCT_StyleDefinition()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleDefinition should validate: %s", err)
	}
}

func TestCT_StyleDefinitionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleDefinition()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleDefinition()
	xml.Unmarshal(buf, v2)
}
