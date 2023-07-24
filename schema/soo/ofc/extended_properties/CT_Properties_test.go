package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

func TestCT_PropertiesConstructor(t *testing.T) {
	v := extended_properties.NewCT_Properties()
	if v == nil {
		t.Errorf("extended_properties.NewCT_Properties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_Properties should validate: %s", err)
	}
}

func TestCT_PropertiesMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_Properties()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_Properties()
	xml.Unmarshal(buf, v2)
}
