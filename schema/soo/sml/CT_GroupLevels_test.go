package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupLevelsConstructor(t *testing.T) {
	v := sml.NewCT_GroupLevels()
	if v == nil {
		t.Errorf("sml.NewCT_GroupLevels must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupLevels should validate: %s", err)
	}
}

func TestCT_GroupLevelsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupLevels()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupLevels()
	xml.Unmarshal(buf, v2)
}
