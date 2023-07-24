package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DirectionConstructor(t *testing.T) {
	v := diagram.NewCT_Direction()
	if v == nil {
		t.Errorf("diagram.NewCT_Direction must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Direction should validate: %s", err)
	}
}

func TestCT_DirectionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Direction()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Direction()
	xml.Unmarshal(buf, v2)
}
