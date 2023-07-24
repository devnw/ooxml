package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_BlockLevelEltsConstructor(t *testing.T) {
	v := wml.NewEG_BlockLevelElts()
	if v == nil {
		t.Errorf("wml.NewEG_BlockLevelElts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_BlockLevelElts should validate: %s", err)
	}
}

func TestEG_BlockLevelEltsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_BlockLevelElts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_BlockLevelElts()
	xml.Unmarshal(buf, v2)
}
