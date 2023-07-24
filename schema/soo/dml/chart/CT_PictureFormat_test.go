package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PictureFormatConstructor(t *testing.T) {
	v := chart.NewCT_PictureFormat()
	if v == nil {
		t.Errorf("chart.NewCT_PictureFormat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PictureFormat should validate: %s", err)
	}
}

func TestCT_PictureFormatMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PictureFormat()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PictureFormat()
	xml.Unmarshal(buf, v2)
}
