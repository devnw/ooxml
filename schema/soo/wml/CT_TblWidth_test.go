package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblWidthConstructor(t *testing.T) {
	v := wml.NewCT_TblWidth()
	if v == nil {
		t.Errorf("wml.NewCT_TblWidth must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblWidth should validate: %s", err)
	}
}

func TestCT_TblWidthMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblWidth()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblWidth()
	xml.Unmarshal(buf, v2)
}
