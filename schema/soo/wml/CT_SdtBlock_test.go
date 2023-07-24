package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtBlockConstructor(t *testing.T) {
	v := wml.NewCT_SdtBlock()
	if v == nil {
		t.Errorf("wml.NewCT_SdtBlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtBlock should validate: %s", err)
	}
}

func TestCT_SdtBlockMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtBlock()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtBlock()
	xml.Unmarshal(buf, v2)
}
