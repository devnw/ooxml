package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocDefaultsConstructor(t *testing.T) {
	v := wml.NewCT_DocDefaults()
	if v == nil {
		t.Errorf("wml.NewCT_DocDefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocDefaults should validate: %s", err)
	}
}

func TestCT_DocDefaultsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocDefaults()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocDefaults()
	xml.Unmarshal(buf, v2)
}
