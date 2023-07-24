package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ForEachConstructor(t *testing.T) {
	v := diagram.NewCT_ForEach()
	if v == nil {
		t.Errorf("diagram.NewCT_ForEach must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ForEach should validate: %s", err)
	}
}

func TestCT_ForEachMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ForEach()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ForEach()
	xml.Unmarshal(buf, v2)
}
