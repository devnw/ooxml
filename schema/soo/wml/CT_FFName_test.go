package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFNameConstructor(t *testing.T) {
	v := wml.NewCT_FFName()
	if v == nil {
		t.Errorf("wml.NewCT_FFName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFName should validate: %s", err)
	}
}

func TestCT_FFNameMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFName()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFName()
	xml.Unmarshal(buf, v2)
}
