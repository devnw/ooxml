package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblGridChangeConstructor(t *testing.T) {
	v := wml.NewCT_TblGridChange()
	if v == nil {
		t.Errorf("wml.NewCT_TblGridChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblGridChange should validate: %s", err)
	}
}

func TestCT_TblGridChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblGridChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblGridChange()
	xml.Unmarshal(buf, v2)
}
