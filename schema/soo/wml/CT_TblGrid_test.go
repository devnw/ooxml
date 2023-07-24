package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblGridConstructor(t *testing.T) {
	v := wml.NewCT_TblGrid()
	if v == nil {
		t.Errorf("wml.NewCT_TblGrid must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblGrid should validate: %s", err)
	}
}

func TestCT_TblGridMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblGrid()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblGrid()
	xml.Unmarshal(buf, v2)
}
