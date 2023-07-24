package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BodyConstructor(t *testing.T) {
	v := wml.NewCT_Body()
	if v == nil {
		t.Errorf("wml.NewCT_Body must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Body should validate: %s", err)
	}
}

func TestCT_BodyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Body()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Body()
	xml.Unmarshal(buf, v2)
}
