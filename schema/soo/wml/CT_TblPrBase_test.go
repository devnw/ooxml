package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPrBaseConstructor(t *testing.T) {
	v := wml.NewCT_TblPrBase()
	if v == nil {
		t.Errorf("wml.NewCT_TblPrBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPrBase should validate: %s", err)
	}
}

func TestCT_TblPrBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPrBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPrBase()
	xml.Unmarshal(buf, v2)
}
