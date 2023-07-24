package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

func TestPropertiesConstructor(t *testing.T) {
	v := extended_properties.NewProperties()
	if v == nil {
		t.Errorf("extended_properties.NewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.Properties should validate: %s", err)
	}
}

func TestPropertiesMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewProperties()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewProperties()
	xml.Unmarshal(buf, v2)
}
