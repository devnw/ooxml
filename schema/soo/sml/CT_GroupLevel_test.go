package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupLevelConstructor(t *testing.T) {
	v := sml.NewCT_GroupLevel()
	if v == nil {
		t.Errorf("sml.NewCT_GroupLevel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupLevel should validate: %s", err)
	}
}

func TestCT_GroupLevelMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupLevel()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupLevel()
	xml.Unmarshal(buf, v2)
}
