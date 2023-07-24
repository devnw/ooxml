package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_RangeMarkupElementsConstructor(t *testing.T) {
	v := wml.NewEG_RangeMarkupElements()
	if v == nil {
		t.Errorf("wml.NewEG_RangeMarkupElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RangeMarkupElements should validate: %s", err)
	}
}

func TestEG_RangeMarkupElementsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RangeMarkupElements()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RangeMarkupElements()
	xml.Unmarshal(buf, v2)
}
