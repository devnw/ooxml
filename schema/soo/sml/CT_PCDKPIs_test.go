package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PCDKPIsConstructor(t *testing.T) {
	v := sml.NewCT_PCDKPIs()
	if v == nil {
		t.Errorf("sml.NewCT_PCDKPIs must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PCDKPIs should validate: %s", err)
	}
}

func TestCT_PCDKPIsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PCDKPIs()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PCDKPIs()
	xml.Unmarshal(buf, v2)
}
