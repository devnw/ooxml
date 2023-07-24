package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtPrConstructor(t *testing.T) {
	v := wml.NewCT_SdtPr()
	if v == nil {
		t.Errorf("wml.NewCT_SdtPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtPr should validate: %s", err)
	}
}

func TestCT_SdtPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtPr()
	xml.Unmarshal(buf, v2)
}
