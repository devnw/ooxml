package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_VMergeConstructor(t *testing.T) {
	v := wml.NewCT_VMerge()
	if v == nil {
		t.Errorf("wml.NewCT_VMerge must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_VMerge should validate: %s", err)
	}
}

func TestCT_VMergeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_VMerge()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_VMerge()
	xml.Unmarshal(buf, v2)
}
