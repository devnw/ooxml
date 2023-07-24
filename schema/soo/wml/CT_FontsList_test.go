package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontsListConstructor(t *testing.T) {
	v := wml.NewCT_FontsList()
	if v == nil {
		t.Errorf("wml.NewCT_FontsList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FontsList should validate: %s", err)
	}
}

func TestCT_FontsListMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FontsList()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FontsList()
	xml.Unmarshal(buf, v2)
}
