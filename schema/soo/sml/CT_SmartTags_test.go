package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SmartTagsConstructor(t *testing.T) {
	v := sml.NewCT_SmartTags()
	if v == nil {
		t.Errorf("sml.NewCT_SmartTags must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SmartTags should validate: %s", err)
	}
}

func TestCT_SmartTagsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SmartTags()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SmartTags()
	xml.Unmarshal(buf, v2)
}
