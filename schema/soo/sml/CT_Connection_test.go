package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConnectionConstructor(t *testing.T) {
	v := sml.NewCT_Connection()
	if v == nil {
		t.Errorf("sml.NewCT_Connection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Connection should validate: %s", err)
	}
}

func TestCT_ConnectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Connection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Connection()
	xml.Unmarshal(buf, v2)
}
