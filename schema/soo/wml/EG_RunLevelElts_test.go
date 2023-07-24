package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RunLevelEltsConstructor(t *testing.T) {
	v := wml.NewEG_RunLevelElts()
	if v == nil {
		t.Errorf("wml.NewEG_RunLevelElts must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RunLevelElts should validate: %s", err)
	}
}

func TestEG_RunLevelEltsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RunLevelElts()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RunLevelElts()
	xml.Unmarshal(buf, v2)
}
