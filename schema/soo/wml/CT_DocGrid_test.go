package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocGridConstructor(t *testing.T) {
	v := wml.NewCT_DocGrid()
	if v == nil {
		t.Errorf("wml.NewCT_DocGrid must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocGrid should validate: %s", err)
	}
}

func TestCT_DocGridMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocGrid()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocGrid()
	xml.Unmarshal(buf, v2)
}
