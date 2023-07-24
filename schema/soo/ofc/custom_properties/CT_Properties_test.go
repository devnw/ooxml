package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/custom_properties"
)

func TestCT_PropertiesConstructor(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	if v == nil {
		t.Errorf("custom_properties.NewCT_Properties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.CT_Properties should validate: %s", err)
	}
}

func TestCT_PropertiesMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewCT_Properties()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewCT_Properties()
	xml.Unmarshal(buf, v2)
}
