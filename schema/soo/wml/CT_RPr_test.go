package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RPrConstructor(t *testing.T) {
	v := wml.NewCT_RPr()
	if v == nil {
		t.Errorf("wml.NewCT_RPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RPr should validate: %s", err)
	}
}

func TestCT_RPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RPr()
	xml.Unmarshal(buf, v2)
}
