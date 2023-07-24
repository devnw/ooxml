package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EmConstructor(t *testing.T) {
	v := wml.NewCT_Em()
	if v == nil {
		t.Errorf("wml.NewCT_Em must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Em should validate: %s", err)
	}
}

func TestCT_EmMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Em()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Em()
	xml.Unmarshal(buf, v2)
}
