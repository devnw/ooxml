package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/custom_properties"
)

func TestCT_PropertyConstructor(t *testing.T) {
	v := custom_properties.NewCT_Property()
	if v == nil {
		t.Errorf("custom_properties.NewCT_Property must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.CT_Property should validate: %s", err)
	}
}

func TestCT_PropertyMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewCT_Property()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewCT_Property()
	xml.Unmarshal(buf, v2)
}
