package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DocRsidsConstructor(t *testing.T) {
	v := wml.NewCT_DocRsids()
	if v == nil {
		t.Errorf("wml.NewCT_DocRsids must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DocRsids should validate: %s", err)
	}
}

func TestCT_DocRsidsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DocRsids()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DocRsids()
	xml.Unmarshal(buf, v2)
}
