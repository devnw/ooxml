package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_AxSharedConstructor(t *testing.T) {
	v := chart.NewEG_AxShared()
	if v == nil {
		t.Errorf("chart.NewEG_AxShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_AxShared should validate: %s", err)
	}
}

func TestEG_AxSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_AxShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_AxShared()
	xml.Unmarshal(buf, v2)
}
