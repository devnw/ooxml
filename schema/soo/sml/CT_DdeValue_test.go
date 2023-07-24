package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DdeValueConstructor(t *testing.T) {
	v := sml.NewCT_DdeValue()
	if v == nil {
		t.Errorf("sml.NewCT_DdeValue must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DdeValue should validate: %s", err)
	}
}

func TestCT_DdeValueMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DdeValue()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DdeValue()
	xml.Unmarshal(buf, v2)
}
