package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AltChunkPrConstructor(t *testing.T) {
	v := wml.NewCT_AltChunkPr()
	if v == nil {
		t.Errorf("wml.NewCT_AltChunkPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_AltChunkPr should validate: %s", err)
	}
}

func TestCT_AltChunkPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_AltChunkPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_AltChunkPr()
	xml.Unmarshal(buf, v2)
}
