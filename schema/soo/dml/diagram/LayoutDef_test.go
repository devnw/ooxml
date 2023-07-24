package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestLayoutDefConstructor(t *testing.T) {
	v := diagram.NewLayoutDef()
	if v == nil {
		t.Errorf("diagram.NewLayoutDef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.LayoutDef should validate: %s", err)
	}
}

func TestLayoutDefMarshalUnmarshal(t *testing.T) {
	v := diagram.NewLayoutDef()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewLayoutDef()
	xml.Unmarshal(buf, v2)
}
