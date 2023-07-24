package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MergeCellsConstructor(t *testing.T) {
	v := sml.NewCT_MergeCells()
	if v == nil {
		t.Errorf("sml.NewCT_MergeCells must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MergeCells should validate: %s", err)
	}
}

func TestCT_MergeCellsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MergeCells()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MergeCells()
	xml.Unmarshal(buf, v2)
}
