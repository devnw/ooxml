package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblOverlapConstructor(t *testing.T) {
	v := wml.NewCT_TblOverlap()
	if v == nil {
		t.Errorf("wml.NewCT_TblOverlap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblOverlap should validate: %s", err)
	}
}

func TestCT_TblOverlapMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblOverlap()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblOverlap()
	xml.Unmarshal(buf, v2)
}
