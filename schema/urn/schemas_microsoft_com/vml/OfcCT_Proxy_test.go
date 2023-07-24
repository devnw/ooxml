package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcCT_ProxyConstructor(t *testing.T) {
	v := vml.NewOfcCT_Proxy()
	if v == nil {
		t.Errorf("vml.NewOfcCT_Proxy must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcCT_Proxy should validate: %s", err)
	}
}

func TestOfcCT_ProxyMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcCT_Proxy()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcCT_Proxy()
	xml.Unmarshal(buf, v2)
}
