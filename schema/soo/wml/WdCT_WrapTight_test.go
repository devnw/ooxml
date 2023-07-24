package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapTightConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapTight()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapTight must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapTight should validate: %s", err)
	}
}

func TestWdCT_WrapTightMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapTight()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapTight()
	xml.Unmarshal(buf, v2)
}
