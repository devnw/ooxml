package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PictureStackUnitConstructor(t *testing.T) {
	v := chart.NewCT_PictureStackUnit()
	if v == nil {
		t.Errorf("chart.NewCT_PictureStackUnit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PictureStackUnit should validate: %s", err)
	}
}

func TestCT_PictureStackUnitMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PictureStackUnit()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PictureStackUnit()
	xml.Unmarshal(buf, v2)
}
