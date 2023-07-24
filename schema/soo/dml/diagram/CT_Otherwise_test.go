package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_OtherwiseConstructor(t *testing.T) {
	v := diagram.NewCT_Otherwise()
	if v == nil {
		t.Errorf("diagram.NewCT_Otherwise must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Otherwise should validate: %s", err)
	}
}

func TestCT_OtherwiseMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Otherwise()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Otherwise()
	xml.Unmarshal(buf, v2)
}
