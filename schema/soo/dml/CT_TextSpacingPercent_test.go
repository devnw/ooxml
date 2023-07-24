package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextSpacingPercentConstructor(t *testing.T) {
	v := dml.NewCT_TextSpacingPercent()
	if v == nil {
		t.Errorf("dml.NewCT_TextSpacingPercent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextSpacingPercent should validate: %s", err)
	}
}

func TestCT_TextSpacingPercentMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextSpacingPercent()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextSpacingPercent()
	xml.Unmarshal(buf, v2)
}
