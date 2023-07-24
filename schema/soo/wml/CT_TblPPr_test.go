package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblPPrConstructor(t *testing.T) {
	v := wml.NewCT_TblPPr()
	if v == nil {
		t.Errorf("wml.NewCT_TblPPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblPPr should validate: %s", err)
	}
}

func TestCT_TblPPrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblPPr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblPPr()
	xml.Unmarshal(buf, v2)
}
