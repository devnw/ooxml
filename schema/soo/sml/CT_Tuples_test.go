package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TuplesConstructor(t *testing.T) {
	v := sml.NewCT_Tuples()
	if v == nil {
		t.Errorf("sml.NewCT_Tuples must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Tuples should validate: %s", err)
	}
}

func TestCT_TuplesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Tuples()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Tuples()
	xml.Unmarshal(buf, v2)
}
