package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestHdrConstructor(t *testing.T) {
	v := wml.NewHdr()
	if v == nil {
		t.Errorf("wml.NewHdr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Hdr should validate: %s", err)
	}
}

func TestHdrMarshalUnmarshal(t *testing.T) {
	v := wml.NewHdr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewHdr()
	xml.Unmarshal(buf, v2)
}
