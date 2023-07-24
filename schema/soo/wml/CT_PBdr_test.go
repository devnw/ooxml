package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PBdrConstructor(t *testing.T) {
	v := wml.NewCT_PBdr()
	if v == nil {
		t.Errorf("wml.NewCT_PBdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PBdr should validate: %s", err)
	}
}

func TestCT_PBdrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PBdr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PBdr()
	xml.Unmarshal(buf, v2)
}
