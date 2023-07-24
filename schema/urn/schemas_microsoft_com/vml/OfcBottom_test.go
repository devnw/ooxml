package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcBottomConstructor(t *testing.T) {
	v := vml.NewOfcBottom()
	if v == nil {
		t.Errorf("vml.NewOfcBottom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcBottom should validate: %s", err)
	}
}

func TestOfcBottomMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcBottom()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcBottom()
	xml.Unmarshal(buf, v2)
}
