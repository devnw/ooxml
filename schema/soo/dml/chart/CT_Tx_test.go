package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_TxConstructor(t *testing.T) {
	v := chart.NewCT_Tx()
	if v == nil {
		t.Errorf("chart.NewCT_Tx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Tx should validate: %s", err)
	}
}

func TestCT_TxMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Tx()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Tx()
	xml.Unmarshal(buf, v2)
}
