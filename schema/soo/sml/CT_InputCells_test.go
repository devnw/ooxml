package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_InputCellsConstructor(t *testing.T) {
	v := sml.NewCT_InputCells()
	if v == nil {
		t.Errorf("sml.NewCT_InputCells must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_InputCells should validate: %s", err)
	}
}

func TestCT_InputCellsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_InputCells()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_InputCells()
	xml.Unmarshal(buf, v2)
}
