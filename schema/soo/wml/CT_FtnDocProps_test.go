package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FtnDocPropsConstructor(t *testing.T) {
	v := wml.NewCT_FtnDocProps()
	if v == nil {
		t.Errorf("wml.NewCT_FtnDocProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FtnDocProps should validate: %s", err)
	}
}

func TestCT_FtnDocPropsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FtnDocProps()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FtnDocProps()
	xml.Unmarshal(buf, v2)
}
