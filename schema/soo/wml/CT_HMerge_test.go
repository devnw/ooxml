package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HMergeConstructor(t *testing.T) {
	v := wml.NewCT_HMerge()
	if v == nil {
		t.Errorf("wml.NewCT_HMerge must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_HMerge should validate: %s", err)
	}
}

func TestCT_HMergeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_HMerge()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_HMerge()
	xml.Unmarshal(buf, v2)
}
