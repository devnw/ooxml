package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ColorsConstructor(t *testing.T) {
	v := diagram.NewCT_Colors()
	if v == nil {
		t.Errorf("diagram.NewCT_Colors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Colors should validate: %s", err)
	}
}

func TestCT_ColorsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Colors()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Colors()
	xml.Unmarshal(buf, v2)
}
