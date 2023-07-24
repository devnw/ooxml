package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryCacheConstructor(t *testing.T) {
	v := sml.NewCT_QueryCache()
	if v == nil {
		t.Errorf("sml.NewCT_QueryCache must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_QueryCache should validate: %s", err)
	}
}

func TestCT_QueryCacheMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_QueryCache()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_QueryCache()
	xml.Unmarshal(buf, v2)
}
