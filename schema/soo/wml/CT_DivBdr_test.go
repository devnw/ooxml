package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DivBdrConstructor(t *testing.T) {
	v := wml.NewCT_DivBdr()
	if v == nil {
		t.Errorf("wml.NewCT_DivBdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DivBdr should validate: %s", err)
	}
}

func TestCT_DivBdrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DivBdr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DivBdr()
	xml.Unmarshal(buf, v2)
}
