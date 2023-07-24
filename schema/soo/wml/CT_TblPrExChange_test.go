package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPrExChangeConstructor(t *testing.T) {
	v := wml.NewCT_TblPrExChange()
	if v == nil {
		t.Errorf("wml.NewCT_TblPrExChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPrExChange should validate: %s", err)
	}
}

func TestCT_TblPrExChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPrExChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPrExChange()
	xml.Unmarshal(buf, v2)
}
