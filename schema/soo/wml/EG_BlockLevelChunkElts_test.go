package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_BlockLevelChunkEltsConstructor(t *testing.T) {
	v := wml.NewEG_BlockLevelChunkElts()
	if v == nil {
		t.Errorf("wml.NewEG_BlockLevelChunkElts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_BlockLevelChunkElts should validate: %s", err)
	}
}

func TestEG_BlockLevelChunkEltsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_BlockLevelChunkElts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_BlockLevelChunkElts()
	xml.Unmarshal(buf, v2)
}
