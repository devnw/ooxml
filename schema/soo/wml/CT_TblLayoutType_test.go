package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblLayoutTypeConstructor(t *testing.T) {
	v := wml.NewCT_TblLayoutType()
	if v == nil {
		t.Errorf("wml.NewCT_TblLayoutType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblLayoutType should validate: %s", err)
	}
}

func TestCT_TblLayoutTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblLayoutType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblLayoutType()
	xml.Unmarshal(buf, v2)
}
