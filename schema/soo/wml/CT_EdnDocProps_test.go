package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EdnDocPropsConstructor(t *testing.T) {
	v := wml.NewCT_EdnDocProps()
	if v == nil {
		t.Errorf("wml.NewCT_EdnDocProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_EdnDocProps should validate: %s", err)
	}
}

func TestCT_EdnDocPropsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_EdnDocProps()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_EdnDocProps()
	xml.Unmarshal(buf, v2)
}
