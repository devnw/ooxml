package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CharacterSpacingConstructor(t *testing.T) {
	v := wml.NewCT_CharacterSpacing()
	if v == nil {
		t.Errorf("wml.NewCT_CharacterSpacing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_CharacterSpacing should validate: %s", err)
	}
}

func TestCT_CharacterSpacingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_CharacterSpacing()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_CharacterSpacing()
	xml.Unmarshal(buf, v2)
}
