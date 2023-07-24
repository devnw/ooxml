package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ConnectionsConstructor(t *testing.T) {
	v := sml.NewCT_Connections()
	if v == nil {
		t.Errorf("sml.NewCT_Connections must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Connections should validate: %s", err)
	}
}

func TestCT_ConnectionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Connections()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Connections()
	xml.Unmarshal(buf, v2)
}
