package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtContentRowConstructor(t *testing.T) {
	v := wml.NewCT_SdtContentRow()
	if v == nil {
		t.Errorf("wml.NewCT_SdtContentRow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtContentRow should validate: %s", err)
	}
}

func TestCT_SdtContentRowMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtContentRow()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtContentRow()
	xml.Unmarshal(buf, v2)
}
