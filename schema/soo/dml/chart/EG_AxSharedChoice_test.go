package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestEG_AxSharedChoiceConstructor(t *testing.T) {
	v := chart.NewEG_AxSharedChoice()
	if v == nil {
		t.Errorf("chart.NewEG_AxSharedChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_AxSharedChoice should validate: %s", err)
	}
}

func TestEG_AxSharedChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_AxSharedChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_AxSharedChoice()
	xml.Unmarshal(buf, v2)
}
