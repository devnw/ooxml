package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ChartFormatsConstructor(t *testing.T) {
	v := sml.NewCT_ChartFormats()
	if v == nil {
		t.Errorf("sml.NewCT_ChartFormats must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ChartFormats should validate: %s", err)
	}
}

func TestCT_ChartFormatsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ChartFormats()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ChartFormats()
	xml.Unmarshal(buf, v2)
}
