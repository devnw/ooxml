package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_LangConstructor(t *testing.T) {
	v := wml.NewCT_Lang()
	if v == nil {
		t.Errorf("wml.NewCT_Lang must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Lang should validate: %s", err)
	}
}

func TestCT_LangMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Lang()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Lang()
	xml.Unmarshal(buf, v2)
}
