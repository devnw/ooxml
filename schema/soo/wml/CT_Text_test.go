package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TextConstructor(t *testing.T) {
	v := wml.NewCT_Text()
	if v == nil {
		t.Errorf("wml.NewCT_Text must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Text should validate: %s", err)
	}
}

func TestCT_TextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Text()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Text()
	xml.Unmarshal(buf, v2)
}
