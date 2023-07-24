package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RecordConstructor(t *testing.T) {
	v := sml.NewCT_Record()
	if v == nil {
		t.Errorf("sml.NewCT_Record must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Record should validate: %s", err)
	}
}

func TestCT_RecordMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Record()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Record()
	xml.Unmarshal(buf, v2)
}
