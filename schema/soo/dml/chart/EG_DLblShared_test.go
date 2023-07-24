package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_DLblSharedConstructor(t *testing.T) {
	v := chart.NewEG_DLblShared()
	if v == nil {
		t.Errorf("chart.NewEG_DLblShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_DLblShared should validate: %s", err)
	}
}

func TestEG_DLblSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_DLblShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_DLblShared()
	xml.Unmarshal(buf, v2)
}
