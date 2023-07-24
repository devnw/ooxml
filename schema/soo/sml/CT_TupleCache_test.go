package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TupleCacheConstructor(t *testing.T) {
	v := sml.NewCT_TupleCache()
	if v == nil {
		t.Errorf("sml.NewCT_TupleCache must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TupleCache should validate: %s", err)
	}
}

func TestCT_TupleCacheMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TupleCache()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TupleCache()
	xml.Unmarshal(buf, v2)
}
