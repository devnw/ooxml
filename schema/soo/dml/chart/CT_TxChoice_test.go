package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TxChoiceConstructor(t *testing.T) {
	v := chart.NewCT_TxChoice()
	if v == nil {
		t.Errorf("chart.NewCT_TxChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TxChoice should validate: %s", err)
	}
}

func TestCT_TxChoiceMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TxChoice()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TxChoice()
	xml.Unmarshal(buf, v2)
}
