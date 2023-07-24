package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestCT_PrintSettingsConstructor(t *testing.T) {
	v := chart.NewCT_PrintSettings()
	if v == nil {
		t.Errorf("chart.NewCT_PrintSettings must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PrintSettings should validate: %s", err)
	}
}

func TestCT_PrintSettingsMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PrintSettings()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PrintSettings()
	xml.Unmarshal(buf, v2)
}
