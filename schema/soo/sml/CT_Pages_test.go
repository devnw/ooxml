package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PagesConstructor(t *testing.T) {
	v := sml.NewCT_Pages()
	if v == nil {
		t.Errorf("sml.NewCT_Pages must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Pages should validate: %s", err)
	}
}

func TestCT_PagesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Pages()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Pages()
	xml.Unmarshal(buf, v2)
}
