package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FitTextConstructor(t *testing.T) {
	v := wml.NewCT_FitText()
	if v == nil {
		t.Errorf("wml.NewCT_FitText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FitText should validate: %s", err)
	}
}

func TestCT_FitTextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FitText()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FitText()
	xml.Unmarshal(buf, v2)
}
