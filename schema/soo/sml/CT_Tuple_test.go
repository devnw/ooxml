package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TupleConstructor(t *testing.T) {
	v := sml.NewCT_Tuple()
	if v == nil {
		t.Errorf("sml.NewCT_Tuple must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Tuple should validate: %s", err)
	}
}

func TestCT_TupleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Tuple()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Tuple()
	xml.Unmarshal(buf, v2)
}
