package vml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestAG_OfficeShapeAttributesConstructor(t *testing.T) {
	v := vml.NewAG_OfficeShapeAttributes()
	if v == nil {
		t.Errorf("vml.NewAG_OfficeShapeAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_OfficeShapeAttributes should validate: %s", err)
	}
}

func TestAG_OfficeShapeAttributesMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_OfficeShapeAttributes()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_OfficeShapeAttributes()
	xml.Unmarshal(buf, v2)
}
