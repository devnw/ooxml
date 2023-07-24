package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PageSetUpPrConstructor(t *testing.T) {
	v := sml.NewCT_PageSetUpPr()
	if v == nil {
		t.Errorf("sml.NewCT_PageSetUpPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageSetUpPr should validate: %s", err)
	}
}

func TestCT_PageSetUpPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageSetUpPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageSetUpPr()
	xml.Unmarshal(buf, v2)
}
