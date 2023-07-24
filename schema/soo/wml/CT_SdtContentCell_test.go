package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtContentCellConstructor(t *testing.T) {
	v := wml.NewCT_SdtContentCell()
	if v == nil {
		t.Errorf("wml.NewCT_SdtContentCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtContentCell should validate: %s", err)
	}
}

func TestCT_SdtContentCellMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtContentCell()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtContentCell()
	xml.Unmarshal(buf, v2)
}
