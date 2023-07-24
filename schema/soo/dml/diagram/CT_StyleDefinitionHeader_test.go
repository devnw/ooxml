package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_StyleDefinitionHeaderConstructor(t *testing.T) {
	v := diagram.NewCT_StyleDefinitionHeader()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleDefinitionHeader must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleDefinitionHeader should validate: %s", err)
	}
}

func TestCT_StyleDefinitionHeaderMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleDefinitionHeader()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleDefinitionHeader()
	xml.Unmarshal(buf, v2)
}
