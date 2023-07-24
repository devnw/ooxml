package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellConstructor(t *testing.T) {
	v := sml.NewCT_Cell()
	if v == nil {
		t.Errorf("sml.NewCT_Cell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Cell should validate: %s", err)
	}
}

func TestCT_CellMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Cell()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Cell()
	xml.Unmarshal(buf, v2)
}
