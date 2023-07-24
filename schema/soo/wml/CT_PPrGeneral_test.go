package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PPrGeneralConstructor(t *testing.T) {
	v := wml.NewCT_PPrGeneral()
	if v == nil {
		t.Errorf("wml.NewCT_PPrGeneral must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PPrGeneral should validate: %s", err)
	}
}

func TestCT_PPrGeneralMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PPrGeneral()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PPrGeneral()
	xml.Unmarshal(buf, v2)
}
