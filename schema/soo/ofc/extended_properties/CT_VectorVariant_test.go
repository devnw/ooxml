package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

func TestCT_VectorVariantConstructor(t *testing.T) {
	v := extended_properties.NewCT_VectorVariant()
	if v == nil {
		t.Errorf("extended_properties.NewCT_VectorVariant must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_VectorVariant should validate: %s", err)
	}
}

func TestCT_VectorVariantMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_VectorVariant()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_VectorVariant()
	xml.Unmarshal(buf, v2)
}
