package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetConstructor(t *testing.T) {
	v := sml.NewCT_Sheet()
	if v == nil {
		t.Errorf("sml.NewCT_Sheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Sheet should validate: %s", err)
	}
}

func TestCT_SheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Sheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Sheet()
	xml.Unmarshal(buf, v2)
}
