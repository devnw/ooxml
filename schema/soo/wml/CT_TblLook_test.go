package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblLookConstructor(t *testing.T) {
	v := wml.NewCT_TblLook()
	if v == nil {
		t.Errorf("wml.NewCT_TblLook must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblLook should validate: %s", err)
	}
}

func TestCT_TblLookMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblLook()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblLook()
	xml.Unmarshal(buf, v2)
}
