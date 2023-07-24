package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DbPrConstructor(t *testing.T) {
	v := sml.NewCT_DbPr()
	if v == nil {
		t.Errorf("sml.NewCT_DbPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DbPr should validate: %s", err)
	}
}

func TestCT_DbPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DbPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DbPr()
	xml.Unmarshal(buf, v2)
}
