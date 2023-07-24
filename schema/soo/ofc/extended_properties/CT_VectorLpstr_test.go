package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

func TestCT_VectorLpstrConstructor(t *testing.T) {
	v := extended_properties.NewCT_VectorLpstr()
	if v == nil {
		t.Errorf("extended_properties.NewCT_VectorLpstr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_VectorLpstr should validate: %s", err)
	}
}

func TestCT_VectorLpstrMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_VectorLpstr()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_VectorLpstr()
	xml.Unmarshal(buf, v2)
}
