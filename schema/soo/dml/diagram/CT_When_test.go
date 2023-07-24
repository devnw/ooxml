package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_WhenConstructor(t *testing.T) {
	v := diagram.NewCT_When()
	if v == nil {
		t.Errorf("diagram.NewCT_When must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_When should validate: %s", err)
	}
}

func TestCT_WhenMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_When()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_When()
	xml.Unmarshal(buf, v2)
}
