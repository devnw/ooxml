package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TablesConstructor(t *testing.T) {
	v := sml.NewCT_Tables()
	if v == nil {
		t.Errorf("sml.NewCT_Tables must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Tables should validate: %s", err)
	}
}

func TestCT_TablesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Tables()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Tables()
	xml.Unmarshal(buf, v2)
}
