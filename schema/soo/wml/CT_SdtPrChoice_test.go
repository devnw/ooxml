package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtPrChoiceConstructor(t *testing.T) {
	v := wml.NewCT_SdtPrChoice()
	if v == nil {
		t.Errorf("wml.NewCT_SdtPrChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtPrChoice should validate: %s", err)
	}
}

func TestCT_SdtPrChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtPrChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtPrChoice()
	xml.Unmarshal(buf, v2)
}
