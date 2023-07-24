package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_SerTxConstructor(t *testing.T) {
	v := chart.NewCT_SerTx()
	if v == nil {
		t.Errorf("chart.NewCT_SerTx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_SerTx should validate: %s", err)
	}
}

func TestCT_SerTxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_SerTx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_SerTx()
	xml.Unmarshal(buf, v2)
}
