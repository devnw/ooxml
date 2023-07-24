package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_SerSharedConstructor(t *testing.T) {
	v := chart.NewEG_SerShared()
	if v == nil {
		t.Errorf("chart.NewEG_SerShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_SerShared should validate: %s", err)
	}
}

func TestEG_SerSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_SerShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_SerShared()
	xml.Unmarshal(buf, v2)
}
