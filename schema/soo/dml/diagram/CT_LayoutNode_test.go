package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_LayoutNodeConstructor(t *testing.T) {
	v := diagram.NewCT_LayoutNode()
	if v == nil {
		t.Errorf("diagram.NewCT_LayoutNode must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_LayoutNode should validate: %s", err)
	}
}

func TestCT_LayoutNodeMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_LayoutNode()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_LayoutNode()
	xml.Unmarshal(buf, v2)
}
