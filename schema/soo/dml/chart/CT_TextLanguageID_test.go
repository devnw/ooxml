package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TextLanguageIDConstructor(t *testing.T) {
	v := chart.NewCT_TextLanguageID()
	if v == nil {
		t.Errorf("chart.NewCT_TextLanguageID must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TextLanguageID should validate: %s", err)
	}
}

func TestCT_TextLanguageIDMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TextLanguageID()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TextLanguageID()
	xml.Unmarshal(buf, v2)
}
