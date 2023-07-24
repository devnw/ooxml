package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_CellMarkupElementsConstructor(t *testing.T) {
	v := wml.NewEG_CellMarkupElements()
	if v == nil {
		t.Errorf("wml.NewEG_CellMarkupElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_CellMarkupElements should validate: %s", err)
	}
}

func TestEG_CellMarkupElementsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_CellMarkupElements()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_CellMarkupElements()
	xml.Unmarshal(buf, v2)
}
