package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PictureOptionsConstructor(t *testing.T) {
	v := chart.NewCT_PictureOptions()
	if v == nil {
		t.Errorf("chart.NewCT_PictureOptions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PictureOptions should validate: %s", err)
	}
}

func TestCT_PictureOptionsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PictureOptions()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PictureOptions()
	xml.Unmarshal(buf, v2)
}
