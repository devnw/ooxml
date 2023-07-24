package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestBackgroundConstructor(t *testing.T) {
	v := vml.NewBackground()
	if v == nil {
		t.Errorf("vml.NewBackground must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Background should validate: %s", err)
	}
}

func TestBackgroundMarshalUnmarshal(t *testing.T) {
	v := vml.NewBackground()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewBackground()
	xml.Unmarshal(buf, v2)
}
