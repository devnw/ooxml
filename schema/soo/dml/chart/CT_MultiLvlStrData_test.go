package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_MultiLvlStrDataConstructor(t *testing.T) {
	v := chart.NewCT_MultiLvlStrData()
	if v == nil {
		t.Errorf("chart.NewCT_MultiLvlStrData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_MultiLvlStrData should validate: %s", err)
	}
}

func TestCT_MultiLvlStrDataMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_MultiLvlStrData()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_MultiLvlStrData()
	xml.Unmarshal(buf, v2)
}
