package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BackgroundConstructor(t *testing.T) {
	v := wml.NewCT_Background()
	if v == nil {
		t.Errorf("wml.NewCT_Background must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Background should validate: %s", err)
	}
}

func TestCT_BackgroundMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Background()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Background()
	xml.Unmarshal(buf, v2)
}
