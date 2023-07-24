package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtEndPrConstructor(t *testing.T) {
	v := wml.NewCT_SdtEndPr()
	if v == nil {
		t.Errorf("wml.NewCT_SdtEndPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtEndPr should validate: %s", err)
	}
}

func TestCT_SdtEndPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtEndPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtEndPr()
	xml.Unmarshal(buf, v2)
}
