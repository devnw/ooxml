package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LvlConstructor(t *testing.T) {
	v := wml.NewCT_Lvl()
	if v == nil {
		t.Errorf("wml.NewCT_Lvl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Lvl should validate: %s", err)
	}
}

func TestCT_LvlMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Lvl()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Lvl()
	xml.Unmarshal(buf, v2)
}
