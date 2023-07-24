package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_OptimizeForBrowserConstructor(t *testing.T) {
	v := wml.NewCT_OptimizeForBrowser()
	if v == nil {
		t.Errorf("wml.NewCT_OptimizeForBrowser must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_OptimizeForBrowser should validate: %s", err)
	}
}

func TestCT_OptimizeForBrowserMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_OptimizeForBrowser()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_OptimizeForBrowser()
	xml.Unmarshal(buf, v2)
}
