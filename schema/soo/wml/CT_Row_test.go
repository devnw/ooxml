package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RowConstructor(t *testing.T) {
	v := wml.NewCT_Row()
	if v == nil {
		t.Errorf("wml.NewCT_Row must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Row should validate: %s", err)
	}
}

func TestCT_RowMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Row()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Row()
	xml.Unmarshal(buf, v2)
}
