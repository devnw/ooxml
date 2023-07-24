package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtDateConstructor(t *testing.T) {
	v := wml.NewCT_SdtDate()
	if v == nil {
		t.Errorf("wml.NewCT_SdtDate must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtDate should validate: %s", err)
	}
}

func TestCT_SdtDateMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtDate()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtDate()
	xml.Unmarshal(buf, v2)
}
