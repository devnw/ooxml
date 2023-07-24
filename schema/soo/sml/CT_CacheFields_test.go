package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CacheFieldsConstructor(t *testing.T) {
	v := sml.NewCT_CacheFields()
	if v == nil {
		t.Errorf("sml.NewCT_CacheFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CacheFields should validate: %s", err)
	}
}

func TestCT_CacheFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CacheFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CacheFields()
	xml.Unmarshal(buf, v2)
}
