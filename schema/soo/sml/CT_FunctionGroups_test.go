package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FunctionGroupsConstructor(t *testing.T) {
	v := sml.NewCT_FunctionGroups()
	if v == nil {
		t.Errorf("sml.NewCT_FunctionGroups must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FunctionGroups should validate: %s", err)
	}
}

func TestCT_FunctionGroupsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FunctionGroups()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FunctionGroups()
	xml.Unmarshal(buf, v2)
}
