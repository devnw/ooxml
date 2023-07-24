package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PageSetupConstructor(t *testing.T) {
	v := sml.NewCT_PageSetup()
	if v == nil {
		t.Errorf("sml.NewCT_PageSetup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageSetup should validate: %s", err)
	}
}

func TestCT_PageSetupMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageSetup()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageSetup()
	xml.Unmarshal(buf, v2)
}
