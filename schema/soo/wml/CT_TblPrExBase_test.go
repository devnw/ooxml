package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPrExBaseConstructor(t *testing.T) {
	v := wml.NewCT_TblPrExBase()
	if v == nil {
		t.Errorf("wml.NewCT_TblPrExBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPrExBase should validate: %s", err)
	}
}

func TestCT_TblPrExBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPrExBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPrExBase()
	xml.Unmarshal(buf, v2)
}
