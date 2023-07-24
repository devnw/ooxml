package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtCellConstructor(t *testing.T) {
	v := wml.NewCT_SdtCell()
	if v == nil {
		t.Errorf("wml.NewCT_SdtCell must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtCell should validate: %s", err)
	}
}

func TestCT_SdtCellMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtCell()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtCell()
	xml.Unmarshal(buf, v2)
}
