package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellWatchConstructor(t *testing.T) {
	v := sml.NewCT_CellWatch()
	if v == nil {
		t.Errorf("sml.NewCT_CellWatch must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellWatch should validate: %s", err)
	}
}

func TestCT_CellWatchMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellWatch()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellWatch()
	xml.Unmarshal(buf, v2)
}
