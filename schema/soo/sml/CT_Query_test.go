package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryConstructor(t *testing.T) {
	v := sml.NewCT_Query()
	if v == nil {
		t.Errorf("sml.NewCT_Query must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Query should validate: %s", err)
	}
}

func TestCT_QueryMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Query()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Query()
	xml.Unmarshal(buf, v2)
}
