package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EdnPropsConstructor(t *testing.T) {
	v := wml.NewCT_EdnProps()
	if v == nil {
		t.Errorf("wml.NewCT_EdnProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_EdnProps should validate: %s", err)
	}
}

func TestCT_EdnPropsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_EdnProps()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_EdnProps()
	xml.Unmarshal(buf, v2)
}
