package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtRowConstructor(t *testing.T) {
	v := wml.NewCT_SdtRow()
	if v == nil {
		t.Errorf("wml.NewCT_SdtRow must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtRow should validate: %s", err)
	}
}

func TestCT_SdtRowMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtRow()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtRow()
	xml.Unmarshal(buf, v2)
}
