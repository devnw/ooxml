package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TblGridColConstructor(t *testing.T) {
	v := wml.NewCT_TblGridCol()
	if v == nil {
		t.Errorf("wml.NewCT_TblGridCol must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TblGridCol should validate: %s", err)
	}
}

func TestCT_TblGridColMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TblGridCol()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TblGridCol()
	xml.Unmarshal(buf, v2)
}
