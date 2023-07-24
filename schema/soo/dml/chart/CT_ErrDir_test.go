package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_ErrDirConstructor(t *testing.T) {
	v := chart.NewCT_ErrDir()
	if v == nil {
		t.Errorf("chart.NewCT_ErrDir must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_ErrDir should validate: %s", err)
	}
}

func TestCT_ErrDirMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_ErrDir()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_ErrDir()
	xml.Unmarshal(buf, v2)
}
