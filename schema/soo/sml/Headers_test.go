package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestHeadersConstructor(t *testing.T) {
	v := sml.NewHeaders()
	if v == nil {
		t.Errorf("sml.NewHeaders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Headers should validate: %s", err)
	}
}

func TestHeadersMarshalUnmarshal(t *testing.T) {
	v := sml.NewHeaders()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewHeaders()
	xml.Unmarshal(buf, v2)
}
