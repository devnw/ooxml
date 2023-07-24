package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PTabConstructor(t *testing.T) {
	v := wml.NewCT_PTab()
	if v == nil {
		t.Errorf("wml.NewCT_PTab must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PTab should validate: %s", err)
	}
}

func TestCT_PTabMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PTab()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PTab()
	xml.Unmarshal(buf, v2)
}
