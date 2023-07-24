package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtDocPartConstructor(t *testing.T) {
	v := wml.NewCT_SdtDocPart()
	if v == nil {
		t.Errorf("wml.NewCT_SdtDocPart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtDocPart should validate: %s", err)
	}
}

func TestCT_SdtDocPartMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtDocPart()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtDocPart()
	xml.Unmarshal(buf, v2)
}
