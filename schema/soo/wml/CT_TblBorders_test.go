package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblBordersConstructor(t *testing.T) {
	v := wml.NewCT_TblBorders()
	if v == nil {
		t.Errorf("wml.NewCT_TblBorders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblBorders should validate: %s", err)
	}
}

func TestCT_TblBordersMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblBorders()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblBorders()
	xml.Unmarshal(buf, v2)
}
