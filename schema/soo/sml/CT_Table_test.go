package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableConstructor(t *testing.T) {
	v := sml.NewCT_Table()
	if v == nil {
		t.Errorf("sml.NewCT_Table must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Table should validate: %s", err)
	}
}

func TestCT_TableMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Table()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Table()
	xml.Unmarshal(buf, v2)
}
