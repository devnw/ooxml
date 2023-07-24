package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtTextConstructor(t *testing.T) {
	v := wml.NewCT_SdtText()
	if v == nil {
		t.Errorf("wml.NewCT_SdtText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtText should validate: %s", err)
	}
}

func TestCT_SdtTextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtText()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtText()
	xml.Unmarshal(buf, v2)
}
