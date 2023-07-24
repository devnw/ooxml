package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LvlLegacyConstructor(t *testing.T) {
	v := wml.NewCT_LvlLegacy()
	if v == nil {
		t.Errorf("wml.NewCT_LvlLegacy must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_LvlLegacy should validate: %s", err)
	}
}

func TestCT_LvlLegacyMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_LvlLegacy()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_LvlLegacy()
	xml.Unmarshal(buf, v2)
}
