package core_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/metadata/core_properties"
)

func TestCorePropertiesConstructor(t *testing.T) {
	v := core_properties.NewCoreProperties()
	if v == nil {
		t.Errorf("core_properties.NewCoreProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CoreProperties should validate: %s", err)
	}
}

func TestCorePropertiesMarshalUnmarshal(t *testing.T) {
	v := core_properties.NewCoreProperties()
	buf, _ := xml.Marshal(v)
	v2 := core_properties.NewCoreProperties()
	xml.Unmarshal(buf, v2)
}
