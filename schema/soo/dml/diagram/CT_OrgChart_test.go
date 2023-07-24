package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_OrgChartConstructor(t *testing.T) {
	v := diagram.NewCT_OrgChart()
	if v == nil {
		t.Errorf("diagram.NewCT_OrgChart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_OrgChart should validate: %s", err)
	}
}

func TestCT_OrgChartMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_OrgChart()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_OrgChart()
	xml.Unmarshal(buf, v2)
}
