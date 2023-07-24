package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcLockConstructor(t *testing.T) {
	v := vml.NewOfcLock()
	if v == nil {
		t.Errorf("vml.NewOfcLock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcLock should validate: %s", err)
	}
}

func TestOfcLockMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcLock()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcLock()
	xml.Unmarshal(buf, v2)
}
