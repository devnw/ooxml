package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextLineBreakConstructor(t *testing.T) {
	v := dml.NewCT_TextLineBreak()
	if v == nil {
		t.Errorf("dml.NewCT_TextLineBreak must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextLineBreak should validate: %s", err)
	}
}

func TestCT_TextLineBreakMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextLineBreak()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextLineBreak()
	xml.Unmarshal(buf, v2)
}
