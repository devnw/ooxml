package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextTabStopListConstructor(t *testing.T) {
	v := dml.NewCT_TextTabStopList()
	if v == nil {
		t.Errorf("dml.NewCT_TextTabStopList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextTabStopList should validate: %s", err)
	}
}

func TestCT_TextTabStopListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextTabStopList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextTabStopList()
	xml.Unmarshal(buf, v2)
}
