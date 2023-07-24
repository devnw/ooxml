package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalCellConstructor(t *testing.T) {
	v := sml.NewCT_ExternalCell()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalCell should validate: %s", err)
	}
}

func TestCT_ExternalCellMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalCell()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalCell()
	xml.Unmarshal(buf, v2)
}
