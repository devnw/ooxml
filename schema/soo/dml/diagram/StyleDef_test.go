package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestStyleDefConstructor(t *testing.T) {
	v := diagram.NewStyleDef()
	if v == nil {
		t.Errorf("diagram.NewStyleDef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.StyleDef should validate: %s", err)
	}
}

func TestStyleDefMarshalUnmarshal(t *testing.T) {
	v := diagram.NewStyleDef()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewStyleDef()
	xml.Unmarshal(buf, v2)
}
