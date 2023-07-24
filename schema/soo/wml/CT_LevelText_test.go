package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LevelTextConstructor(t *testing.T) {
	v := wml.NewCT_LevelText()
	if v == nil {
		t.Errorf("wml.NewCT_LevelText must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LevelText should validate: %s", err)
	}
}

func TestCT_LevelTextMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LevelText()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LevelText()
	xml.Unmarshal(buf, v2)
}
