package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcColumnConstructor(t *testing.T) {
	v := vml.NewOfcColumn()
	if v == nil {
		t.Errorf("vml.NewOfcColumn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcColumn should validate: %s", err)
	}
}

func TestOfcColumnMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcColumn()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcColumn()
	xml.Unmarshal(buf, v2)
}
