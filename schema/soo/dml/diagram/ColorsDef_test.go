package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestColorsDefConstructor(t *testing.T) {
	v := diagram.NewColorsDef()
	if v == nil {
		t.Errorf("diagram.NewColorsDef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.ColorsDef should validate: %s", err)
	}
}

func TestColorsDefMarshalUnmarshal(t *testing.T) {
	v := diagram.NewColorsDef()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewColorsDef()
	xml.Unmarshal(buf, v2)
}
