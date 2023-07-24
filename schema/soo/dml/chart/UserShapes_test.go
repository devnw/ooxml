package chart_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chart"
)

func TestUserShapesConstructor(t *testing.T) {
	v := chart.NewUserShapes()
	if v == nil {
		t.Errorf("chart.NewUserShapes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.UserShapes should validate: %s", err)
	}
}

func TestUserShapesMarshalUnmarshal(t *testing.T) {
	v := chart.NewUserShapes()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewUserShapes()
	xml.Unmarshal(buf, v2)
}
