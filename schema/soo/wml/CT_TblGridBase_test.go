package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblGridBaseConstructor(t *testing.T) {
	v := wml.NewCT_TblGridBase()
	if v == nil {
		t.Errorf("wml.NewCT_TblGridBase must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblGridBase should validate: %s", err)
	}
}

func TestCT_TblGridBaseMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblGridBase()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblGridBase()
	xml.Unmarshal(buf, v2)
}
