package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotCacheConstructor(t *testing.T) {
	v := sml.NewCT_PivotCache()
	if v == nil {
		t.Errorf("sml.NewCT_PivotCache must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotCache should validate: %s", err)
	}
}

func TestCT_PivotCacheMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotCache()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotCache()
	xml.Unmarshal(buf, v2)
}
