package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_StyleDefinitionHeaderLstConstructor(t *testing.T) {
	v := diagram.NewCT_StyleDefinitionHeaderLst()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleDefinitionHeaderLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleDefinitionHeaderLst should validate: %s", err)
	}
}

func TestCT_StyleDefinitionHeaderLstMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleDefinitionHeaderLst()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleDefinitionHeaderLst()
	xml.Unmarshal(buf, v2)
}
