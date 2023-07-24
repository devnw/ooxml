package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ServerFormatsConstructor(t *testing.T) {
	v := sml.NewCT_ServerFormats()
	if v == nil {
		t.Errorf("sml.NewCT_ServerFormats must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ServerFormats should validate: %s", err)
	}
}

func TestCT_ServerFormatsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ServerFormats()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ServerFormats()
	xml.Unmarshal(buf, v2)
}
