package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_CrossBetweenConstructor(t *testing.T) {
	v := chart.NewCT_CrossBetween()
	if v == nil {
		t.Errorf("chart.NewCT_CrossBetween must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_CrossBetween should validate: %s", err)
	}
}

func TestCT_CrossBetweenMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_CrossBetween()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_CrossBetween()
	xml.Unmarshal(buf, v2)
}
