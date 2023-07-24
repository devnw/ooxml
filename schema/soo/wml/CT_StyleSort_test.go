package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_StyleSortConstructor(t *testing.T) {
	v := wml.NewCT_StyleSort()
	if v == nil {
		t.Errorf("wml.NewCT_StyleSort must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_StyleSort should validate: %s", err)
	}
}

func TestCT_StyleSortMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_StyleSort()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_StyleSort()
	xml.Unmarshal(buf, v2)
}
