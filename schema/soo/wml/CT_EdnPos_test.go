package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_EdnPosConstructor(t *testing.T) {
	v := wml.NewCT_EdnPos()
	if v == nil {
		t.Errorf("wml.NewCT_EdnPos must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_EdnPos should validate: %s", err)
	}
}

func TestCT_EdnPosMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_EdnPos()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_EdnPos()
	xml.Unmarshal(buf, v2)
}
