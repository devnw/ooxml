package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FontSigConstructor(t *testing.T) {
	v := wml.NewCT_FontSig()
	if v == nil {
		t.Errorf("wml.NewCT_FontSig must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FontSig should validate: %s", err)
	}
}

func TestCT_FontSigMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FontSig()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FontSig()
	xml.Unmarshal(buf, v2)
}
