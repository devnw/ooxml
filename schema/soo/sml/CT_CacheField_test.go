package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CacheFieldConstructor(t *testing.T) {
	v := sml.NewCT_CacheField()
	if v == nil {
		t.Errorf("sml.NewCT_CacheField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CacheField should validate: %s", err)
	}
}

func TestCT_CacheFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CacheField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CacheField()
	xml.Unmarshal(buf, v2)
}
