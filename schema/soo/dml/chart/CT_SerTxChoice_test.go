package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SerTxChoiceConstructor(t *testing.T) {
	v := chart.NewCT_SerTxChoice()
	if v == nil {
		t.Errorf("chart.NewCT_SerTxChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SerTxChoice should validate: %s", err)
	}
}

func TestCT_SerTxChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SerTxChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SerTxChoice()
	xml.Unmarshal(buf, v2)
}
