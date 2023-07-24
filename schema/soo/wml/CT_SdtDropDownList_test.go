package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtDropDownListConstructor(t *testing.T) {
	v := wml.NewCT_SdtDropDownList()
	if v == nil {
		t.Errorf("wml.NewCT_SdtDropDownList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtDropDownList should validate: %s", err)
	}
}

func TestCT_SdtDropDownListMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtDropDownList()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtDropDownList()
	xml.Unmarshal(buf, v2)
}
