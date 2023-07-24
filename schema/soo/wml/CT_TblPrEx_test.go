package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPrExConstructor(t *testing.T) {
	v := wml.NewCT_TblPrEx()
	if v == nil {
		t.Errorf("wml.NewCT_TblPrEx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPrEx should validate: %s", err)
	}
}

func TestCT_TblPrExMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPrEx()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPrEx()
	xml.Unmarshal(buf, v2)
}
