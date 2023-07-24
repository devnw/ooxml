package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblConstructor(t *testing.T) {
	v := wml.NewCT_Tbl()
	if v == nil {
		t.Errorf("wml.NewCT_Tbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Tbl should validate: %s", err)
	}
}

func TestCT_TblMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Tbl()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Tbl()
	xml.Unmarshal(buf, v2)
}
