package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LevelSuffixConstructor(t *testing.T) {
	v := wml.NewCT_LevelSuffix()
	if v == nil {
		t.Errorf("wml.NewCT_LevelSuffix must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LevelSuffix should validate: %s", err)
	}
}

func TestCT_LevelSuffixMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LevelSuffix()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LevelSuffix()
	xml.Unmarshal(buf, v2)
}
