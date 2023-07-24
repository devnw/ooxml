package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotCachesConstructor(t *testing.T) {
	v := sml.NewCT_PivotCaches()
	if v == nil {
		t.Errorf("sml.NewCT_PivotCaches must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotCaches should validate: %s", err)
	}
}

func TestCT_PivotCachesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotCaches()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotCaches()
	xml.Unmarshal(buf, v2)
}
