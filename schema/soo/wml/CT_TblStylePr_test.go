package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblStylePrConstructor(t *testing.T) {
	v := wml.NewCT_TblStylePr()
	if v == nil {
		t.Errorf("wml.NewCT_TblStylePr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblStylePr should validate: %s", err)
	}
}

func TestCT_TblStylePrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblStylePr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblStylePr()
	xml.Unmarshal(buf, v2)
}
