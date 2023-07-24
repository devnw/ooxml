package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOvalConstructor(t *testing.T) {
	v := vml.NewOval()
	if v == nil {
		t.Errorf("vml.NewOval must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Oval should validate: %s", err)
	}
}

func TestOvalMarshalUnmarshal(t *testing.T) {
	v := vml.NewOval()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOval()
	xml.Unmarshal(buf, v2)
}
