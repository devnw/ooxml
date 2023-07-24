package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SortConditionConstructor(t *testing.T) {
	v := sml.NewCT_SortCondition()
	if v == nil {
		t.Errorf("sml.NewCT_SortCondition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SortCondition should validate: %s", err)
	}
}

func TestCT_SortConditionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SortCondition()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SortCondition()
	xml.Unmarshal(buf, v2)
}
