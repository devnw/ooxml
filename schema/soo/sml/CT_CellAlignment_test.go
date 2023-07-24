package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellAlignmentConstructor(t *testing.T) {
	v := sml.NewCT_CellAlignment()
	if v == nil {
		t.Errorf("sml.NewCT_CellAlignment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellAlignment should validate: %s", err)
	}
}

func TestCT_CellAlignmentMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellAlignment()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellAlignment()
	xml.Unmarshal(buf, v2)
}
