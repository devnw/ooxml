package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DdeLinkConstructor(t *testing.T) {
	v := sml.NewCT_DdeLink()
	if v == nil {
		t.Errorf("sml.NewCT_DdeLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DdeLink should validate: %s", err)
	}
}

func TestCT_DdeLinkMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DdeLink()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DdeLink()
	xml.Unmarshal(buf, v2)
}
