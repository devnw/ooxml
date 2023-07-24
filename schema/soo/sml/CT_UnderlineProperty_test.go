package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_UnderlinePropertyConstructor(t *testing.T) {
	v := sml.NewCT_UnderlineProperty()
	if v == nil {
		t.Errorf("sml.NewCT_UnderlineProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_UnderlineProperty should validate: %s", err)
	}
}

func TestCT_UnderlinePropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_UnderlineProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_UnderlineProperty()
	xml.Unmarshal(buf, v2)
}
