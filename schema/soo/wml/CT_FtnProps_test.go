package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FtnPropsConstructor(t *testing.T) {
	v := wml.NewCT_FtnProps()
	if v == nil {
		t.Errorf("wml.NewCT_FtnProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnProps should validate: %s", err)
	}
}

func TestCT_FtnPropsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnProps()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnProps()
	xml.Unmarshal(buf, v2)
}
