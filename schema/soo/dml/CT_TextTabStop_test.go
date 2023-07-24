package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextTabStopConstructor(t *testing.T) {
	v := dml.NewCT_TextTabStop()
	if v == nil {
		t.Errorf("dml.NewCT_TextTabStop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextTabStop should validate: %s", err)
	}
}

func TestCT_TextTabStopMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextTabStop()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextTabStop()
	xml.Unmarshal(buf, v2)
}
