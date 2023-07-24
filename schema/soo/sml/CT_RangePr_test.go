package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RangePrConstructor(t *testing.T) {
	v := sml.NewCT_RangePr()
	if v == nil {
		t.Errorf("sml.NewCT_RangePr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RangePr should validate: %s", err)
	}
}

func TestCT_RangePrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RangePr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RangePr()
	xml.Unmarshal(buf, v2)
}
