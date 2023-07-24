package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_DashStopConstructor(t *testing.T) {
	v := dml.NewCT_DashStop()
	if v == nil {
		t.Errorf("dml.NewCT_DashStop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_DashStop should validate: %s", err)
	}
}

func TestCT_DashStopMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_DashStop()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_DashStop()
	xml.Unmarshal(buf, v2)
}
