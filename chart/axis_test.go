package chart_test

import (
	"testing"

	"go.devnw.com/ooxml/chart"
)

func TestNullAxis(t *testing.T) {
	if chart.NullAxis.AxisID() != 0 {
		t.Errorf("expected null axis to have ID 0, go %d", chart.NullAxis.AxisID())
	}
}
