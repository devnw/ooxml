package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DdeValuesConstructor(t *testing.T) {
	v := sml.NewCT_DdeValues()
	if v == nil {
		t.Errorf("sml.NewCT_DdeValues must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DdeValues should validate: %s", err)
	}
}

func TestCT_DdeValuesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DdeValues()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DdeValues()
	xml.Unmarshal(buf, v2)
}
