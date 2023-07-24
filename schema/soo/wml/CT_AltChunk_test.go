package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AltChunkConstructor(t *testing.T) {
	v := wml.NewCT_AltChunk()
	if v == nil {
		t.Errorf("wml.NewCT_AltChunk must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_AltChunk should validate: %s", err)
	}
}

func TestCT_AltChunkMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_AltChunk()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_AltChunk()
	xml.Unmarshal(buf, v2)
}
