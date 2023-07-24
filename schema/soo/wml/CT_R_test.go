package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RConstructor(t *testing.T) {
	v := wml.NewCT_R()
	if v == nil {
		t.Errorf("wml.NewCT_R must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_R should validate: %s", err)
	}
}

func TestCT_RMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_R()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_R()
	xml.Unmarshal(buf, v2)
}
