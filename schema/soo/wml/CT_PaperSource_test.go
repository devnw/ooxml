package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PaperSourceConstructor(t *testing.T) {
	v := wml.NewCT_PaperSource()
	if v == nil {
		t.Errorf("wml.NewCT_PaperSource must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PaperSource should validate: %s", err)
	}
}

func TestCT_PaperSourceMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PaperSource()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PaperSource()
	xml.Unmarshal(buf, v2)
}
