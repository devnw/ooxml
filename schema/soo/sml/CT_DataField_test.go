package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataFieldConstructor(t *testing.T) {
	v := sml.NewCT_DataField()
	if v == nil {
		t.Errorf("sml.NewCT_DataField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataField should validate: %s", err)
	}
}

func TestCT_DataFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataField()
	xml.Unmarshal(buf, v2)
}
