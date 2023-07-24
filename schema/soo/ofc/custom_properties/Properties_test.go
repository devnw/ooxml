package custom_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/custom_properties"
)

func TestPropertiesConstructor(t *testing.T) {
	v := custom_properties.NewProperties()
	if v == nil {
		t.Errorf("custom_properties.NewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed custom_properties.Properties should validate: %s", err)
	}
}

func TestPropertiesMarshalUnmarshal(t *testing.T) {
	v := custom_properties.NewProperties()
	buf, _ := xml.Marshal(v)
	v2 := custom_properties.NewProperties()
	xml.Unmarshal(buf, v2)
}
