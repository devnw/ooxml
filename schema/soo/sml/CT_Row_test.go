package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RowConstructor(t *testing.T) {
	v := sml.NewCT_Row()
	if v == nil {
		t.Errorf("sml.NewCT_Row must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Row should validate: %s", err)
	}
}

func TestCT_RowMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Row()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Row()
	xml.Unmarshal(buf, v2)
}
