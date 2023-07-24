package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ObjectLinkConstructor(t *testing.T) {
	v := wml.NewCT_ObjectLink()
	if v == nil {
		t.Errorf("wml.NewCT_ObjectLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ObjectLink should validate: %s", err)
	}
}

func TestCT_ObjectLinkMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ObjectLink()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ObjectLink()
	xml.Unmarshal(buf, v2)
}
