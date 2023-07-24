package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtContentBlockConstructor(t *testing.T) {
	v := wml.NewCT_SdtContentBlock()
	if v == nil {
		t.Errorf("wml.NewCT_SdtContentBlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtContentBlock should validate: %s", err)
	}
}

func TestCT_SdtContentBlockMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtContentBlock()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtContentBlock()
	xml.Unmarshal(buf, v2)
}
