package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellStyleXfsConstructor(t *testing.T) {
	v := sml.NewCT_CellStyleXfs()
	if v == nil {
		t.Errorf("sml.NewCT_CellStyleXfs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellStyleXfs should validate: %s", err)
	}
}

func TestCT_CellStyleXfsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellStyleXfs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellStyleXfs()
	xml.Unmarshal(buf, v2)
}
