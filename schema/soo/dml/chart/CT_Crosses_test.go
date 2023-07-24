package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_CrossesConstructor(t *testing.T) {
	v := chart.NewCT_Crosses()
	if v == nil {
		t.Errorf("chart.NewCT_Crosses must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_Crosses should validate: %s", err)
	}
}

func TestCT_CrossesMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_Crosses()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_Crosses()
	xml.Unmarshal(buf, v2)
}
