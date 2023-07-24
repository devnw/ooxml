package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_TblPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_TblPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPrChange should validate: %s", err)
	}
}

func TestCT_TblPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPrChange()
	xml.Unmarshal(buf, v2)
}
