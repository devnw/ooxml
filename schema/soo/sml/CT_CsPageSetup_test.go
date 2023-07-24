package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CsPageSetupConstructor(t *testing.T) {
	v := sml.NewCT_CsPageSetup()
	if v == nil {
		t.Errorf("sml.NewCT_CsPageSetup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CsPageSetup should validate: %s", err)
	}
}

func TestCT_CsPageSetupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CsPageSetup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CsPageSetup()
	xml.Unmarshal(buf, v2)
}
