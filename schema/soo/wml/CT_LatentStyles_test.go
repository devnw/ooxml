package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LatentStylesConstructor(t *testing.T) {
	v := wml.NewCT_LatentStyles()
	if v == nil {
		t.Errorf("wml.NewCT_LatentStyles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LatentStyles should validate: %s", err)
	}
}

func TestCT_LatentStylesMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LatentStyles()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LatentStyles()
	xml.Unmarshal(buf, v2)
}
