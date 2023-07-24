package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalBookConstructor(t *testing.T) {
	v := sml.NewCT_ExternalBook()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalBook must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalBook should validate: %s", err)
	}
}

func TestCT_ExternalBookMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalBook()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalBook()
	xml.Unmarshal(buf, v2)
}
