package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableMissingConstructor(t *testing.T) {
	v := sml.NewCT_TableMissing()
	if v == nil {
		t.Errorf("sml.NewCT_TableMissing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableMissing should validate: %s", err)
	}
}

func TestCT_TableMissingMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableMissing()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableMissing()
	xml.Unmarshal(buf, v2)
}
