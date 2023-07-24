package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_MultiLvlStrRefConstructor(t *testing.T) {
	v := chart.NewCT_MultiLvlStrRef()
	if v == nil {
		t.Errorf("chart.NewCT_MultiLvlStrRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_MultiLvlStrRef should validate: %s", err)
	}
}

func TestCT_MultiLvlStrRefMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_MultiLvlStrRef()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_MultiLvlStrRef()
	xml.Unmarshal(buf, v2)
}
